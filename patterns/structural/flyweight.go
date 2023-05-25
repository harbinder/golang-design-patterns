package structural

/*
Refer: https://golangbyexample.com/flyweight-design-pattern-golang/

Flyweight structural design pattern is used when
- when a large number of similar objects need to be created.
- These objects are called flyweight objects and are immutable.


In the flyweight pattern, we store the flyweight objects in the map.
Whenever the other objects which share the flyweight objects are created, then flyweight objects are fetched from the map.

When to Use:
- When the objects have some intrinsic properties which can be shared.
- Use flyweight when a large number of objects needs to be created which can cause memory issue.
  In case figure out all the common or intrinsic state and create flyweight objects for that.
*/

type House struct {
	Name  HouseNameEnum
	Color string
}

type Student struct {
	House      House
	SchoolName string
	Class      int
	Section    string
}

type HouseNameEnum string

const (
	HOUSE_PRITHVI HouseNameEnum = "PRITHVI"
	HOUSE_JAL     HouseNameEnum = "JAL"
	HOUSE_AAKASH  HouseNameEnum = "AKASH"
	HOUSE_AGNI    HouseNameEnum = "AGNI"
)

// Initialise immutable object map of house objects
type HouseFactory struct {
	houseMap map[HouseNameEnum]House
}

func (hm *HouseFactory) InitializeHouseMap(houseName HouseNameEnum) House {
	if _, ok := hm.houseMap[houseName]; ok == true {
		return hm.houseMap[houseName]
	}

	switch houseName {
	case HOUSE_PRITHVI:
		hm.houseMap[houseName] = House{Name: HOUSE_PRITHVI, Color: "Green"}
	case HOUSE_JAL:
		hm.houseMap[houseName] = House{Name: HOUSE_JAL, Color: "White"}
	case HOUSE_AAKASH:
		hm.houseMap[houseName] = House{Name: HOUSE_AAKASH, Color: "Blue"}
	case HOUSE_AGNI:
		hm.houseMap[houseName] = House{Name: HOUSE_AGNI, Color: "Red"}
	}
	return hm.houseMap[houseName]
}
