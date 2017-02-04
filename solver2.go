//A probably horribly inefficient way to solve sudoku puzzles using the power of io
//inspired by: http://elmo.sbs.arizona.edu/sandiway/sudoku/examples.html

//some csv reading tech taken from:
//https://github.com/thbar/golang-playground/blob/master/csv-parsing.go

package main

import(
	"fmt"
	"io"
	"os"
	"encoding/csv"
	"strconv"
	"math/rand"
	"time"
	"math"
	//"reflect"
)


type Vertex struct {
	X, Y int
}

type VertexInfo struct {
	X int
}

type VertexMap map[Vertex]VertexInfo


func Round(f float64) float64 {
	return math.Floor(f + .5)
}

func blocks(slice [][]float64, starti int, stopi int, startj int, stopj int) [][]float64 {
		

	existslice := make([]float64, 0, 10)

	for i := starti; i < stopi; i++ {
		for j := startj; j < stopj; j++{
			if slice[i][j] != 0{
				//fmt.Println(slice[i][j])
				existslice = append(existslice,slice[i][j])
			}
		}
	}

	doesntexistslice := make([]float64, 0, 10)

	counter := 0

	for i := 1; i < 10; i++{

		counter = 0

		for j := 0; j < len(existslice); j++{

			if existslice[j] != float64(i){
				
				counter = counter + 1
		
			}
		}

		if counter == len(existslice){
			doesntexistslice = append(doesntexistslice,float64(i))
		}

	}

	counter = 0

	for i := starti; i < stopi; i++ {
		for j := startj; j < stopj; j++{
			if slice[i][j] == 0{
				slice[i][j] = -(doesntexistslice[counter])
				counter = counter + 1
			}
		}
	}


	//fmt.Println(slice)
	//fmt.Println(existslice)

	return slice
}

func flipper(blockmap VertexMap, slice [][]float64) [][]float64{



// counterintuitive syntax for testing with the vertex struct map structure
	
//	if vm[Vertex{1.5,2}] == (VertexInfo{1}){
//		fmt.Println("hey")
//	}

	dummyslice := make([][]float64,9)
	copy(dummyslice, slice)
	returnslice := make([][]float64,9)
	copy(returnslice, slice)

	var dummynumij float64 = 0
	var dummynumkm float64 = 0
	var fitnessold float64 = 0
	var fitnessnew float64 = 0
	//var epsilon float64 = .0001
	//counter2 := 0
	//stickcounter := 0
	i:=0
	j:=0
	k:=0
	m:=0
	var exitloop float64 = 0 

	//evalnum:=0
	//evaldenom:=0
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

	for exitloop == 0{
		for i = 0; i < 9; i++ {
			for j = 0; j < 9; j++{
				for k = 0; k < 9; k++{
					for m = 0; m < 9; m++{
						if slice[k][m] < 0 && slice[i][j] < 0 && blockmap[Vertex{i,j}] == blockmap[Vertex{k,m}] && !((i==k)&&(j==m)) && !((i==m)&&(j==k)) && exitloop == 0{
							
							s1 = rand.NewSource(time.Now().UnixNano())
    						r1 = rand.New(s1)

							fitnessold = linecheck(dummyslice)

							dummynumij = slice[i][j]
							dummynumkm = slice[k][m]

							dummyslice[i][j] = dummynumkm
							dummyslice[k][m] = dummynumij 


							fitnessnew = linecheck(dummyslice)

							if (fitnessnew<fitnessold) || (r1.Float64()) > .99 {
								//epsilon = epsilon + .001
								//stickcounter = 0
								fmt.Println("values of (i,j) and (k,m): (", i,",",j,") (",k,",",m,")", " blockmaps kj, km: ", blockmap[Vertex{i,j}], blockmap[Vertex{k,m}])
								fmt.Println("fitnessnew: ", fitnessnew)
								fmt.Println("fitnessold: ", fitnessold)
								fmt.Println("random number generate: ", Round(r1.Float64()), r1.Float64())
							//fmt.Println("evalnum: ", evalnum)
								if fitnessnew == 0 {

									fmt.Println("fitnessintheexitif: ", fitnessnew)
									copy(returnslice,dummyslice)
									exitloop = 1

								}
							}else if (fitnessnew > fitnessold) {

								dummyslice[i][j] = dummynumij
								dummyslice[k][m] = dummynumkm
								//stickcounter = stickcounter + 1
								//fmt.Println(dummyslice)
							}

						}

					}
				}
			}
		}
	
	//epsilon = 0
	//stickcounter = stickcounter + 1
	i = 0
	j = 0
	k = 0
	m = 0
	//fmt.Println("counter2hit: ", counter2)
	
	}

	fmt.Println("fitnessafterloop: ", fitnessnew)





	return returnslice

}


//dont need to test line by line because I can just run fitness checker

//func testline(slice []int) string{
	
//	testval := ""

//	var fitcheck int = 0

//	for i := 0; i < 9; i++ {

//		if slice[i]>0{
//			fitcheck = fitcheck + slice[i]
//		}else{
//			fitcheck = fitcheck - slice[i]
//		}

		
//	}

//	if fitcheck == 45{
//		testval = "All fit"
//	}else{
//		testval = "Needs flipping"
//	}

//	return testval
//}



func linecheck(slice [][]float64) float64{

	var checknum float64 = 0

	fmt.Println("linecheck slice: ", slice)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++{

			for k := 0; k < 9; k++{
				if k != i{
					if slice[i][j] == slice[k][j] || slice[i][j] == -slice[k][j]{
						checknum = checknum + 1
					}
				}
			}


			for m := 0; m < 9; m++{
				if m != j{
					if slice[i][j] == slice[i][m] || slice[i][j] == -slice[i][m]{
						checknum = checknum + 1
					}
				}
			}

		}

	}

	return checknum

}



func fitnesschecker(slice [][]float64) float64{

	var fitness float64 = 0

	fitnessi := []float64{}
	//slice = append(slice,original...)

	fitnessj := []float64{}
	//slice = append(slice,original...)

	//appenddummy := make([]int, 1, 1)

	linesum := []float64{0}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++{
			if float64(slice[i][j])<0{
				linesum[0] = linesum[0] - float64(slice[i][j])
			} else{
				linesum[0] = linesum[0] + float64(slice[i][j])
			}
		}

		if (linesum[0]-45)<0{
			linesum[0] = -1 * (linesum[0] - 45)
		} else {
			linesum[0] = (linesum[0] - 45)
		}
		//insertVal = append(insertVal, "lol")
		//appenddummy = append(insertval, )
		linesum[0] = linesum[0]/45
		fitnessj = append(fitnessj, linesum...)
	}

	linesum[0] = 0

	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++{
			if float64(slice[i][j])<0{
				linesum[0] = linesum[0] - float64(slice[i][j])
			} else{
				linesum[0] = linesum[0] + float64(slice[i][j])
			}
		}

	
		if (linesum[0]-45)<0{
			linesum[0] = -1 * (linesum[0] - 45)
		} else {
			linesum[0] = (linesum[0] - 45)
		}
		//appenddummy[0] = linesum
		linesum[0] = linesum[0]/45
		fitnessi = append(fitnessi, linesum...)
	}

	linesum[0] = 0

	for j := 0; j < 9; j++ {
		//fmt.Println(linesum)
		linesum[0] = linesum[0] + fitnessi[j] + fitnessj[j]
	}

	//linesum[0] = 100 - 100 * (linesum[0]/18 - 1)

	//linesum[0] = 100 - linesum[0]

	if 100*(1-linesum[0]/18)<1{
		linesum[0] = -100*(1-linesum[0]/18)
	}else{
		linesum[0] = 100*(1-linesum[0]/18)
	}

	//fmt.Println(fitnessj)
	//fmt.Println(fitnessi)
	//fmt.Println(linesum)

	fitness = linesum[0]

	return fitness

}




func main() {



	file, err := os.Open("unsolved.csv")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var savedunsolved [9][9]float64

	//reader.Comma = ,
	lineCount:=0

	fmt.Println("Here is the unsolved sudoku: ")
	for {
		// read just one record, but we could ReadAll() as well
		record, err := reader.Read()
		// end-of-file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// record is an array of string so is directly printable
		//fmt.Println("Record", lineCount, "is", record, "and has", len(record), "fields")
		// and we can iterate on top of that
		
		

		for i := 0; i < len(record); i++ {
			
			integer, err := strconv.Atoi(record[i])
			
			if err != nil{
				fmt.Println("oops!")
			}
			
			savedunsolved[lineCount][i] = float64(integer)
			fmt.Print(" ", record[i])
		}


		


		//fmt.Println(record[0][0])

		fmt.Println()
		lineCount += 1
	}
	fmt.Println("Watch me solve it!")

	


	sliceofcake:=make([][]float64,9)

	for i:=0; i<9; i++ { 

		sliceofcake[i] = savedunsolved[i][:]
	
	}





	//fmt.Println(reflect.TypeOf(savedunsolved))
    //fmt.Println(reflect.TypeOf(sliceofcake))


	fmt.Println(sliceofcake)

   

   sliceofcake = blocks(sliceofcake,0,3,0,3)
   sliceofcake = blocks(sliceofcake,3,6,0,3)
   sliceofcake = blocks(sliceofcake,6,9,0,3)
   
   sliceofcake = blocks(sliceofcake,0,3,3,6)
   sliceofcake = blocks(sliceofcake,3,6,3,6)
   sliceofcake = blocks(sliceofcake,6,9,3,6)

   sliceofcake = blocks(sliceofcake,0,3,6,9)
   sliceofcake = blocks(sliceofcake,3,6,6,9)
   sliceofcake = blocks(sliceofcake,6,9,6,9)

   fmt.Println(sliceofcake)

	

	blockmap := VertexMap{}


	blockmap[Vertex{0, 0}] = VertexInfo{0} 
	blockmap[Vertex{1, 0}] = VertexInfo{0} 
	blockmap[Vertex{2, 0}] = VertexInfo{0} 

	blockmap[Vertex{0, 1}] = VertexInfo{0} 
	blockmap[Vertex{1, 1}] = VertexInfo{0} 
	blockmap[Vertex{2, 1}] = VertexInfo{0} 

	blockmap[Vertex{0, 2}] = VertexInfo{0} 
	blockmap[Vertex{1, 2}] = VertexInfo{0} 
	blockmap[Vertex{2, 2}] = VertexInfo{0} 



	blockmap[Vertex{0, 3}] = VertexInfo{1} 
	blockmap[Vertex{1, 3}] = VertexInfo{1} 
	blockmap[Vertex{2, 3}] = VertexInfo{1} 

	blockmap[Vertex{0, 4}] = VertexInfo{1} 
	blockmap[Vertex{1, 4}] = VertexInfo{1} 
	blockmap[Vertex{2, 4}] = VertexInfo{1} 

	blockmap[Vertex{0, 5}] = VertexInfo{1} 
	blockmap[Vertex{1, 5}] = VertexInfo{1} 
	blockmap[Vertex{2, 5}] = VertexInfo{1} 



	blockmap[Vertex{0, 6}] = VertexInfo{2} 
	blockmap[Vertex{1, 6}] = VertexInfo{2} 
	blockmap[Vertex{2, 6}] = VertexInfo{2} 

	blockmap[Vertex{0, 7}] = VertexInfo{2} 
	blockmap[Vertex{1, 7}] = VertexInfo{2} 
	blockmap[Vertex{2, 7}] = VertexInfo{2} 

	blockmap[Vertex{0, 8}] = VertexInfo{2} 
	blockmap[Vertex{1, 8}] = VertexInfo{2} 
	blockmap[Vertex{2, 8}] = VertexInfo{2} 


	blockmap[Vertex{3, 0}] = VertexInfo{3} 
	blockmap[Vertex{4, 0}] = VertexInfo{3} 
	blockmap[Vertex{5, 0}] = VertexInfo{3} 

	blockmap[Vertex{3, 1}] = VertexInfo{3} 
	blockmap[Vertex{4, 1}] = VertexInfo{3} 
	blockmap[Vertex{5, 1}] = VertexInfo{3} 

	blockmap[Vertex{3, 2}] = VertexInfo{3} 
	blockmap[Vertex{4, 2}] = VertexInfo{3} 
	blockmap[Vertex{5, 2}] = VertexInfo{3} 



	blockmap[Vertex{3, 3}] = VertexInfo{4} 
	blockmap[Vertex{4, 3}] = VertexInfo{4} 
	blockmap[Vertex{5, 3}] = VertexInfo{4} 

	blockmap[Vertex{3, 4}] = VertexInfo{4} 
	blockmap[Vertex{4, 4}] = VertexInfo{4} 
	blockmap[Vertex{5, 4}] = VertexInfo{4} 

	blockmap[Vertex{3, 5}] = VertexInfo{4} 
	blockmap[Vertex{4, 5}] = VertexInfo{4} 
	blockmap[Vertex{5, 5}] = VertexInfo{4} 



	blockmap[Vertex{3, 6}] = VertexInfo{5} 
	blockmap[Vertex{4, 6}] = VertexInfo{5} 
	blockmap[Vertex{5, 6}] = VertexInfo{5} 

	blockmap[Vertex{3, 7}] = VertexInfo{5} 
	blockmap[Vertex{4, 7}] = VertexInfo{5} 
	blockmap[Vertex{5, 7}] = VertexInfo{5} 

	blockmap[Vertex{3, 8}] = VertexInfo{5} 
	blockmap[Vertex{4, 8}] = VertexInfo{5} 
	blockmap[Vertex{5, 8}] = VertexInfo{5} 



	blockmap[Vertex{6, 0}] = VertexInfo{6} 
	blockmap[Vertex{7, 0}] = VertexInfo{6} 
	blockmap[Vertex{8, 0}] = VertexInfo{6} 

	blockmap[Vertex{6, 1}] = VertexInfo{6} 
	blockmap[Vertex{7, 1}] = VertexInfo{6} 
	blockmap[Vertex{8, 1}] = VertexInfo{6} 

	blockmap[Vertex{6, 2}] = VertexInfo{6} 
	blockmap[Vertex{7, 2}] = VertexInfo{6} 
	blockmap[Vertex{8, 2}] = VertexInfo{6} 



	blockmap[Vertex{6, 3}] = VertexInfo{7} 
	blockmap[Vertex{7, 3}] = VertexInfo{7} 
	blockmap[Vertex{8, 3}] = VertexInfo{7} 

	blockmap[Vertex{6, 4}] = VertexInfo{7} 
	blockmap[Vertex{7, 4}] = VertexInfo{7} 
	blockmap[Vertex{8, 4}] = VertexInfo{7} 

	blockmap[Vertex{6, 5}] = VertexInfo{7} 
	blockmap[Vertex{7, 5}] = VertexInfo{7} 
	blockmap[Vertex{8, 5}] = VertexInfo{7} 



	blockmap[Vertex{6, 6}] = VertexInfo{8} 
	blockmap[Vertex{7, 6}] = VertexInfo{8} 
	blockmap[Vertex{8, 6}] = VertexInfo{8} 

	blockmap[Vertex{6, 7}] = VertexInfo{8} 
	blockmap[Vertex{7, 7}] = VertexInfo{8} 
	blockmap[Vertex{8, 7}] = VertexInfo{8} 

	blockmap[Vertex{6, 8}] = VertexInfo{8} 
	blockmap[Vertex{7, 8}] = VertexInfo{8} 
	blockmap[Vertex{8, 8}] = VertexInfo{8} 


//	linemap := VertexMap{}

//	linemap[Vertex{0, 0}] = VertexInfo{0} 
//	linemap[Vertex{1, 0}] = VertexInfo{0} 
//	linemap[Vertex{2, 0}] = VertexInfo{0} 
//	linemap[Vertex{3, 0}] = VertexInfo{0} 
//	linemap[Vertex{4, 0}] = VertexInfo{0} 
//	linemap[Vertex{5, 0}] = VertexInfo{0} 
//	linemap[Vertex{6, 0}] = VertexInfo{0} 
//	linemap[Vertex{7, 0}] = VertexInfo{0} 
//	linemap[Vertex{8, 0}] = VertexInfo{0} 

//	linemap[Vertex{0, 1}] = VertexInfo{1} 
//	linemap[Vertex{1, 1}] = VertexInfo{1} 
//	linemap[Vertex{2, 1}] = VertexInfo{1} 
//	linemap[Vertex{3, 1}] = VertexInfo{1} 
//	linemap[Vertex{4, 1}] = VertexInfo{1} 
//	linemap[Vertex{5, 1}] = VertexInfo{1} 
//	linemap[Vertex{6, 1}] = VertexInfo{1} 
//	linemap[Vertex{7, 1}] = VertexInfo{1} 
//	linemap[Vertex{8, 1}] = VertexInfo{1} 

//	linemap[Vertex{0, 2}] = VertexInfo{2} 
//	linemap[Vertex{1, 2}] = VertexInfo{2} 
//	linemap[Vertex{2, 2}] = VertexInfo{2} 
//	linemap[Vertex{3, 2}] = VertexInfo{2} 
//	linemap[Vertex{4, 2}] = VertexInfo{2} 
//	linemap[Vertex{5, 2}] = VertexInfo{2} 
//	linemap[Vertex{6, 2}] = VertexInfo{2} 
//	linemap[Vertex{7, 2}] = VertexInfo{2} 
//	linemap[Vertex{8, 2}] = VertexInfo{2} 

//	linemap[Vertex{0, 3}] = VertexInfo{3} 
//	linemap[Vertex{1, 3}] = VertexInfo{3} 
//	linemap[Vertex{2, 3}] = VertexInfo{3} 
//	linemap[Vertex{3, 3}] = VertexInfo{3} 
//	linemap[Vertex{4, 3}] = VertexInfo{3} 
//	linemap[Vertex{5, 3}] = VertexInfo{3} 
//	linemap[Vertex{6, 3}] = VertexInfo{3} 
//	linemap[Vertex{7, 3}] = VertexInfo{3} 
//	linemap[Vertex{8, 3}] = VertexInfo{3} 

//	linemap[Vertex{0, 4}] = VertexInfo{4} 
//	linemap[Vertex{1, 4}] = VertexInfo{4} 
//	linemap[Vertex{2, 4}] = VertexInfo{4} 
//	linemap[Vertex{3, 4}] = VertexInfo{4} 
//	linemap[Vertex{4, 4}] = VertexInfo{4} 
//	linemap[Vertex{5, 4}] = VertexInfo{4} 
//	linemap[Vertex{6, 4}] = VertexInfo{4} 
//	linemap[Vertex{7, 4}] = VertexInfo{4} 
//	linemap[Vertex{8, 4}] = VertexInfo{4} 

//	linemap[Vertex{0, 5}] = VertexInfo{5} 
//	linemap[Vertex{1, 5}] = VertexInfo{5} 
//	linemap[Vertex{2, 5}] = VertexInfo{5} 
//	linemap[Vertex{3, 5}] = VertexInfo{5} 
//	linemap[Vertex{4, 5}] = VertexInfo{5} 
//	linemap[Vertex{5, 5}] = VertexInfo{5} 
//	linemap[Vertex{6, 5}] = VertexInfo{5} 
//	linemap[Vertex{7, 5}] = VertexInfo{5} 
//	linemap[Vertex{8, 5}] = VertexInfo{5} 

//	linemap[Vertex{0, 6}] = VertexInfo{6} 
//	linemap[Vertex{1, 6}] = VertexInfo{6} 
//	linemap[Vertex{2, 6}] = VertexInfo{6} 
//	linemap[Vertex{3, 6}] = VertexInfo{6} 
//	linemap[Vertex{4, 6}] = VertexInfo{6} 
//	linemap[Vertex{5, 6}] = VertexInfo{6} 
//	linemap[Vertex{6, 6}] = VertexInfo{6} 
//	linemap[Vertex{7, 6}] = VertexInfo{6} 
//	linemap[Vertex{8, 6}] = VertexInfo{6} 

//	linemap[Vertex{0, 7}] = VertexInfo{7} 
//	linemap[Vertex{1, 7}] = VertexInfo{7} 
//	linemap[Vertex{2, 7}] = VertexInfo{7} 
//	linemap[Vertex{3, 7}] = VertexInfo{7} 
//	linemap[Vertex{4, 7}] = VertexInfo{7} 
//	linemap[Vertex{5, 7}] = VertexInfo{7} 
//	linemap[Vertex{6, 7}] = VertexInfo{7} 
//	linemap[Vertex{7, 7}] = VertexInfo{7} 
//	linemap[Vertex{8, 7}] = VertexInfo{7} 

//	linemap[Vertex{0, 8}] = VertexInfo{8} 
//	linemap[Vertex{1, 8}] = VertexInfo{8} 
//	linemap[Vertex{2, 8}] = VertexInfo{8} 
//	linemap[Vertex{3, 8}] = VertexInfo{8} 
//	linemap[Vertex{4, 8}] = VertexInfo{8} 
//	linemap[Vertex{5, 8}] = VertexInfo{8} 
//	linemap[Vertex{6, 8}] = VertexInfo{8} 
//	linemap[Vertex{7, 8}] = VertexInfo{8} 
//	linemap[Vertex{8, 8}] = VertexInfo{8} 

 

//	linemap[Vertex{0, 0}] = VertexInfo{9} 
//	linemap[Vertex{0, 1}] = VertexInfo{9} 
//	linemap[Vertex{0, 2}] = VertexInfo{9} 
//	linemap[Vertex{0, 3}] = VertexInfo{9} 
//	linemap[Vertex{0, 4}] = VertexInfo{9} 
//	linemap[Vertex{0, 5}] = VertexInfo{9} 
//	linemap[Vertex{0, 6}] = VertexInfo{9} 
//	linemap[Vertex{0, 7}] = VertexInfo{9} 
//	linemap[Vertex{0, 8}] = VertexInfo{9} 

//	linemap[Vertex{1, 0}] = VertexInfo{10} 
//	linemap[Vertex{1, 1}] = VertexInfo{10} 
//	linemap[Vertex{1, 2}] = VertexInfo{10} 
//	linemap[Vertex{1, 3}] = VertexInfo{10} 
//	linemap[Vertex{1, 4}] = VertexInfo{10} 
//	linemap[Vertex{1, 5}] = VertexInfo{10} 
//	linemap[Vertex{1, 6}] = VertexInfo{10} 
//	linemap[Vertex{1, 7}] = VertexInfo{10} 
//	linemap[Vertex{1, 8}] = VertexInfo{10}

//	linemap[Vertex{2, 0}] = VertexInfo{11} 
//	linemap[Vertex{2, 1}] = VertexInfo{11} 
//	linemap[Vertex{2, 2}] = VertexInfo{11} 
//	linemap[Vertex{2, 3}] = VertexInfo{11} 
//	linemap[Vertex{2, 4}] = VertexInfo{11} 
//	linemap[Vertex{2, 5}] = VertexInfo{11} 
//	linemap[Vertex{2, 6}] = VertexInfo{11} 
//	linemap[Vertex{2, 7}] = VertexInfo{11} 
//	linemap[Vertex{2, 8}] = VertexInfo{11}  

//	linemap[Vertex{3, 0}] = VertexInfo{12} 
//	linemap[Vertex{3, 1}] = VertexInfo{12} 
//	linemap[Vertex{3, 2}] = VertexInfo{12} 
//	linemap[Vertex{3, 3}] = VertexInfo{12} 
//	linemap[Vertex{3, 4}] = VertexInfo{12} 
//	linemap[Vertex{3, 5}] = VertexInfo{12} 
//	linemap[Vertex{3, 6}] = VertexInfo{12} 
//	linemap[Vertex{3, 7}] = VertexInfo{12} 
//	linemap[Vertex{3, 8}] = VertexInfo{12} 

//	linemap[Vertex{4, 0}] = VertexInfo{13} 
//	linemap[Vertex{4, 1}] = VertexInfo{13} 
//	linemap[Vertex{4, 2}] = VertexInfo{13} 
//	linemap[Vertex{4, 3}] = VertexInfo{13} 
//	linemap[Vertex{4, 4}] = VertexInfo{13} 
//	linemap[Vertex{4, 5}] = VertexInfo{13} 
//	linemap[Vertex{4, 6}] = VertexInfo{13} 
//	linemap[Vertex{4, 7}] = VertexInfo{13} 
//	linemap[Vertex{4, 8}] = VertexInfo{13}

//	linemap[Vertex{5, 0}] = VertexInfo{14} 
//	linemap[Vertex{5, 1}] = VertexInfo{14} 
//	linemap[Vertex{5, 2}] = VertexInfo{14} 
//	linemap[Vertex{5, 3}] = VertexInfo{14} 
//	linemap[Vertex{5, 4}] = VertexInfo{14} 
//	linemap[Vertex{5, 5}] = VertexInfo{14} 
//	linemap[Vertex{5, 6}] = VertexInfo{14} 
//	linemap[Vertex{5, 7}] = VertexInfo{14} 
//	linemap[Vertex{5, 8}] = VertexInfo{14}  

//	linemap[Vertex{6, 0}] = VertexInfo{15} 
//	linemap[Vertex{6, 1}] = VertexInfo{15} 
//	linemap[Vertex{6, 2}] = VertexInfo{15} 
//	linemap[Vertex{6, 3}] = VertexInfo{15} 
//	linemap[Vertex{6, 4}] = VertexInfo{15} 
//	linemap[Vertex{6, 5}] = VertexInfo{15} 
//	linemap[Vertex{6, 6}] = VertexInfo{15} 
//	linemap[Vertex{6, 7}] = VertexInfo{15} 
//	linemap[Vertex{6, 8}] = VertexInfo{15}  

//	linemap[Vertex{7, 0}] = VertexInfo{16} 
//	linemap[Vertex{7, 1}] = VertexInfo{16} 
//	linemap[Vertex{7, 2}] = VertexInfo{16} 
//	linemap[Vertex{7, 3}] = VertexInfo{16} 
//	linemap[Vertex{7, 4}] = VertexInfo{16} 
//	linemap[Vertex{7, 5}] = VertexInfo{16} 
//	linemap[Vertex{7, 6}] = VertexInfo{16} 
//	linemap[Vertex{7, 7}] = VertexInfo{16} 
//	linemap[Vertex{7, 8}] = VertexInfo{16}  

//	linemap[Vertex{8, 0}] = VertexInfo{17} 
//	linemap[Vertex{8, 1}] = VertexInfo{17} 
//	linemap[Vertex{8, 2}] = VertexInfo{17} 
//	linemap[Vertex{8, 3}] = VertexInfo{17} 
//	linemap[Vertex{8, 4}] = VertexInfo{17} 
//	linemap[Vertex{8, 5}] = VertexInfo{17} 
//	linemap[Vertex{8, 6}] = VertexInfo{17} 
//	linemap[Vertex{8, 7}] = VertexInfo{17} 
//	linemap[Vertex{8, 8}] = VertexInfo{17}  



	fmt.Println(blockmap)

    var fitness float64 = fitnesschecker(sliceofcake)

    fmt.Println("The percent fitness of our sudoku is: ", fitness, "%")

	// counterintuitive syntax for testing with the vertex struct map structure
	
//	if vm[Vertex{1.5,2}] == (VertexInfo{1}){
//		fmt.Println("hey")
//	}

    fmt.Println(sliceofcake)

	sliceofcake = flipper(blockmap, sliceofcake)

	fmt.Println(sliceofcake)


   	//blockmap := make(map[[2]int]int)
	//blockmap[[0][1]] = 1 

//	fmt.Println(blockmap)

}

