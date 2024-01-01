func sumInt(nums ... int)(int , int){

	var  sum, count int

	for _, num := range nums{
		sum += num
		count ++
	}
	
	return count, sum
}