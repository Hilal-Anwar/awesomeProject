package pkg1

import (
	"fmt"
)
func Heaps_permutation(arr[] int,size int)  {
	if size==1 {
		fmt.Println(arr)
	}else {
		for i := 0; i < size; i++ {
			Heaps_permutation(arr,size-1)
			var tem=0
			if size%2!=0{
              tem=arr[0]
			  arr[0]=arr[size-1]
			}else{
				tem=arr[i]
				arr[i]=arr[size-1]
			}
			arr[size-1]=tem
		}
	}
}