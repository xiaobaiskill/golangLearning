package main


func GetRangeMap(m map[string]string)(string,string){
	for k,v := range m{
		return k,v
	}
	return "",""
}
