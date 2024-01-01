workArray := make([]uint8, 10)
indexArray := make([]uint8, 6)
    
for i := 0; i < 10; i++ {
    fmt.Scan(&workArray[i])
}
for i := 0; i < 6; i++ {
    fmt.Scan(&indexArray[i])   
}
workArray[indexArray[0]], workArray[indexArray[1]] = workArray[indexArray[1]], workArray[indexArray[0]]
workArray[indexArray[2]], workArray[indexArray[3]] = workArray[indexArray[3]], workArray[indexArray[2]]
workArray[indexArray[4]], workArray[indexArray[5]] = workArray[indexArray[5]], workArray[indexArray[4]]
for i := 0; i < 10; i++ {
    fmt.Print(workArray[i], " ")
}
fmt.Printf("type ok")
return