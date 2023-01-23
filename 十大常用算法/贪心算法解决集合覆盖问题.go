package main

import "fmt"

type broadcastMap struct {
	hashMap map[string]bool //用来映射对应地区是否存在的hashmap
}

type radioStation struct {
	coverAreas []string
	coverCount int
}

type radioGroup struct {
	radios []radioStation
}

func getMaxCount(radioStations *radioGroup) int {
	max := radioStations.radios[0].coverCount
	maxIndex := 0
	for i := 0; i < len(radioStations.radios); i++ {
		if radioStations.radios[i].coverCount > max {
			max = radioStations.radios[i].coverCount
			maxIndex = i
		}
	}
	return maxIndex
}
func updateCoverCount(radioStations *radioGroup, areasMap *broadcastMap) {
	for i := 0; i < len(radioStations.radios); i++ {
		newCount := 0
		for j := 0; j < len(radioStations.radios[i].coverAreas); j++ {
			if _, ok := areasMap.hashMap[radioStations.radios[i].coverAreas[j]]; ok {
				newCount++
			}
		}
		radioStations.radios[i].coverCount = newCount
	}
}

func greedyAlgorithm(radioStations radioGroup) radioGroup {
	//先把传进来的arrAreas映射进map
	areasMap := broadcastMap{hashMap: make(map[string]bool)}
	for _, v := range radioStations.radios {
		for _, v2 := range v.coverAreas {
			areasMap.hashMap[v2] = true
		}
	}
	//开始执行贪心
	resultRadioGroup := radioGroup{}

	for i := 0; len(areasMap.hashMap) > 0; i++ {
		//根据getMaxCount
		index := getMaxCount(&radioStations)
		for j := 0; j < len(radioStations.radios[index].coverAreas); j++ {
			//删除map中包含的地区
			delete(areasMap.hashMap, radioStations.radios[index].coverAreas[j])
		}
		//把被选中的电台添加到resultRadioGroup
		resultRadioGroup.radios = append(resultRadioGroup.radios, radioStations.radios[index])
		//更新电台组的最大覆盖数
		updateCoverCount(&radioStations, &areasMap)
	}
	return resultRadioGroup
} //返回一个选择方案

func main() {
	radioAreas := radioGroup{radios: make([]radioStation, 0)}
	k1 := radioStation{coverAreas: []string{"北京", "上海", "天津"}, coverCount: 3}
	k2 := radioStation{coverAreas: []string{"北京", "广州", "深圳"}, coverCount: 3}
	k3 := radioStation{coverAreas: []string{"成都", "上海", "杭州"}, coverCount: 3}
	k4 := radioStation{coverAreas: []string{"天津", "上海"}, coverCount: 2}
	k5 := radioStation{coverAreas: []string{"大连", "杭州"}, coverCount: 2}
	radioAreas.radios = append(radioAreas.radios, k1, k2, k3, k4, k5)
	fmt.Println(greedyAlgorithm(radioAreas))

}
