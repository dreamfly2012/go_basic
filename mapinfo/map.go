package mapinfo

/**
* Map return
*
 */
func GenerateMap() map[string]string {
	mymap := make(map[string]string)
	mymap["name"] = "menghuiguli"
	mymap["age"] = "20"
	mymap["score"] = "100"
	return mymap
}
