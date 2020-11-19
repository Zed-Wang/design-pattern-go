package factory

import (
	"fmt"
	"strings"
)

// 将公共的接口抽出来
type AnimalImpl interface {
	Eat()
	Run()
}

// 创建 dog 和 cat 实现接口
type dog struct {
	name string
}

func (d *dog) Eat() {
	fmt.Println("Dog is eating")
}

func (d *dog) Run() {
	fmt.Println("Dog is running")
}

// 比狗多了age属性
type cat struct {
	dog
	age int
}

func (c *cat) Eat() {
	fmt.Println("Cat is eating")
}

func (c *cat) Run() {
	fmt.Println("Cat is running")
}

// 工厂函数
type AnimalFactory func(conf map[string]interface{}) AnimalImpl

// 这些工厂方法接收同样的参数
func NewDog(conf map[string]interface{}) AnimalImpl {
	fmt.Println("Dog: Create")
	name, ok := conf["Name"]
	if !ok {
		fmt.Println("[Name] has not been set in config map")
		return nil
	}
	return &dog{
		name: name.(string),
	}
}

func NewCat(conf map[string]interface{}) AnimalImpl {
	fmt.Println("Cat: Create")
	age, ok := conf["Age"]
	if !ok {
		fmt.Println("[Age] has not been set in config map")
	}
	return &cat{
		age: age.(int),
	}
}

// 注册工厂
type animalTypeEnum string

const (
	Dog animalTypeEnum = "Dog"
	Cat animalTypeEnum = "Cat"
)

var animalFactories = make(map[animalTypeEnum]AnimalFactory) // 用来存放注册了的工厂函数

func RegisterAnimal(animalType animalTypeEnum, factory AnimalFactory) {
	if factory == nil {
		panic(fmt.Sprintf("Animal factory %s does not exist", string(animalType)))
	}
	_, ok := animalFactories[animalType]
	if ok {
		fmt.Printf("Animal factory %s has been registered already\n", string(animalType))
	} else {
		fmt.Printf("Register animal factory %s\n", string(animalType))
		// 注册工厂函数入 animalFactories
		animalFactories[animalType] = factory
	}
}

func init() {
	RegisterAnimal(Dog, NewDog)
	RegisterAnimal(Cat, NewCat)
}

// 创建工厂
func CreateAnimal(conf map[string]interface{}) AnimalImpl {
	animalType, ok := conf["AnimalType"]
	if !ok {
		fmt.Println("No animal type in config map. Use dog type as default.")
		animalType = Dog
	}
	// 从 animalFactories 中获取工厂函数
	AnimalFactory, ok := animalFactories[animalType.(animalTypeEnum)]
	if !ok {
		availableAnimal := make([]string, len(animalFactories))
		for k, _ := range animalFactories {
			availableAnimal = append(availableAnimal, string(k))
		}
		fmt.Printf("Invalid ops type. Must be one of: %s\n", strings.Join(availableAnimal, ", "))
		return nil
	}
	fmt.Println("Create animal: ", animalType)
	// 创建实例
	return AnimalFactory(conf)
}
