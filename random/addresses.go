package random

import "school/types"

func (r *randomGenerator) getAddress() types.Address {

	r.mu.RLock()

	randDisIdx := GetIntBetween(0, len(r.states)-1)
	randCityIdx := GetIntBetween(0, len(r.states[randDisIdx].Cities)-1)

	city := r.states[randDisIdx].Cities[randCityIdx]
	state := r.states[randDisIdx].Name

	r.mu.RUnlock()

	randHouseNo := GetIntBetween(minHouseNum, maxHouseNum)

	return types.Address{
		HouseNumber: randHouseNo,
		State:       state,
		City:        city,
	}

}

func (r *randomGenerator) getName() string {

	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.names[GetIntBetween(0, len(r.names)-1)]
}
