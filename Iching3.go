package main

import "fmt"
import "math/rand"
import "time"
import "bufio"
import "os"

func main() {
	
	var match bool;
	
	var index int;
	
	var row_value int;
	var coin int;
	
	var hexagram [6]int;
	var hexagram2 [6]int;
	
	var hexagrams [384]int;
	
	var row [4]int = [4]int{7,8,6,9};
	
	var rows [4]string = [4]string{"___","_ _","_x_","_o_"};
	
	var texts[64]string = [64]string{"Force","Field","Sprouting","Enveloping","Attending","Arguing","Leading","Grouping","Small Accumulating","Treading","Pervading","Obstruction","Concording People","Great Possessing","Humbling","Providing-For","Following","Correcting","Nearing","Viewing","Gnawing Bite","Adorning","Stripping","Returning","Without Embroiling","Great Accumulating","Swallowing","Great Exceeding","Gorge","Radiance","Conjoining","Persevering","Retiring","Great Invigorating","Prospering","Darkening of the Light","Dwelling People","Polarising","Limping","Taking-Apart","Diminishing","Augmenting","Displacement","Coupling","Clustering","Ascending","Confining","Welling","Skinning","Holding","Shake","Bound","Infiltrating","Converting the Maiden","Abounding","Sojourning","Ground","Open","Dispersing","Articulating","Center Returning","Small Exceeding","Already Fording","Not Yet Fording"};
	
	var bits[6]int = [6]int{1, 2, 4, 8, 16, 32};

	var shuffle[64]int = [64]int{0, 63, 29, 46, 5, 40, 47, 61,
			4, 8, 7, 56, 16, 2, 55, 59,
			25, 38, 15, 60, 26, 22, 62, 31,
			24, 6, 30, 33, 45, 18, 49, 35,
			48, 3, 58, 23, 20, 10, 53, 43,
			14, 28, 1, 32, 57, 39, 41, 37,
			17, 34, 27, 54, 52, 11, 19, 50,
			36, 9, 44, 13, 12, 51, 21, 42};
	
	
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	
	input := bufio.NewScanner(os.Stdin)
	
	for i:=0;i<64;i++ {
		for j:=0;j<6;j++ {
			hexagrams[i*6+j]=(shuffle[i]&bits[j])/bits[j];
		}
	}

	input.Scan()

	for i:=0;i<6;i++ {
		row_value=0;
		for j:=0;j<3;j++ {
			coin=rng.Intn(2)
			row_value+=coin+2;
		}
		for j:=0;j<4;j++ {
			if row_value==row[j] {
				hexagram[i]=j;
			}
		}
	}
	
	for i:=0;i<6;i++ {
		fmt.Printf("%s\n\r", rows[hexagram[i]]);
	}

	input.Scan()

	for i:=0;i<64;i++ {
		match=true;
		for j:=0;j<6;j++ {
			if row[hexagrams[i*6+j]]!=row[hexagram[j]] {
				if row[hexagrams[i*6+j]]!=row[hexagram[j]]+2 {
					if row[hexagrams[i*6+j]]!=row[hexagram[j]]-2 {
						match=false;
					}
				}
			}
		}
		if match {
			index=i;
		}
	}
	
	fmt.Printf("%d\t%s\n\r", index+1, texts[index]);
	
	input.Scan()
	
	for i:=0;i<6;i++ {
		hexagram2[i]=hexagram[i]%2;
	}
	
	for i:=0;i<6;i++ {
		fmt.Printf("%s\n\r", rows[hexagram2[i]]);
	}
		
	input.Scan()
				
	for i:=0;i<64;i++ {
		match=true;
		for j:=0;j<6;j++ {
			if row[hexagrams[i*6+j]]!=row[hexagram2[j]] {
				if row[hexagrams[i*6+j]]!=row[hexagram2[j]]+2 {
					if row[hexagrams[i*6+j]]!=row[hexagram2[j]]-2 {
						match=false;
					}
				}
			}
		}
		if match {
			index=i;
		}
	}
	
	fmt.Printf("%d\t%s\n\r", index+1, texts[index]);

	input.Scan()
	
	return;

}


/*
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define FALSE 0
#define TRUE -1

int main(){

	int i,j;

	int coin;
	int row_value;

	int match;
	int index;

	int hexagram[6];
	int hexagram2[6];

	int hexagrams[384];

	int row[] = {7, 8, 6, 9};
	
	char *row$[] = {"___", "_ _", "_x_", "_o_"};
	
	char *text$[] = {"Force","Field","Sprouting","Enveloping","Attending","Arguing","Leading","Grouping","Small Accumulating","Treading","Pervading","Obstruction","Concording People","Great Possessing","Humbling","Providing-For","Following","Correcting","Nearing","Viewing","Gnawing Bite","Adorning","Stripping","Returning","Without Embroiling","Great Accumulating","Swallowing","Great Exceeding","Gorge","Radiance","Conjoining","Persevering","Retiring","Great Invigorating","Prospering","Darkening of the Light","Dwelling People","Polarising","Limping","Taking-Apart","Diminishing","Augmenting","Displacement","Coupling","Clustering","Ascending","Confining","Welling","Skinning","Holding","Shake","Bound","Infiltrating","Converting the Maiden","Abounding","Sojourning","Ground","Open","Dispersing","Articulating","Center Returning","Small Exceeding","Already Fording","Not Yet Fording"};
	
	int bits[] = {1, 2, 4, 8, 16, 32};

	int shuffle[] = {0, 63, 29, 46, 5, 40, 47, 61,
			4, 8, 7, 56, 16, 2, 55, 59,
			25, 38, 15, 60, 26, 22, 62, 31,
			24, 6, 30, 33, 45, 18, 49, 35,
			48, 3, 58, 23, 20, 10, 53, 43,
			14, 28, 1, 32, 57, 39, 41, 37,
			17, 34, 27, 54, 52, 11, 19, 50,
			36, 9, 44, 13, 12, 51, 21, 42};
	
   	srand((unsigned)time(NULL));
	
	for (i=0;i<64;i++)
		for (j=0;j<6;j++)
			hexagrams[i*6+j]=(shuffle[i]&bits[j])/bits[j];
	
	getchar();

	for (i=0;i<6;i++) {
		row_value=0;
		for (j=0;j<3;j++) {
			coin=(int)((double)rand()*(double)2/(double)RAND_MAX);
			row_value+=coin+2;
		}
		for (j=0;j<4;j++)
			if (row_value==row[j])
				hexagram[i]=j;
	}
	
	for (i=0;i<6;i++)
		printf("%s\n\r", row$[hexagram[i]]);

	getchar();

	for (i=0;i<64;i++) {
		match=TRUE;
		for (j=0;j<6;j++) {
			if (row[hexagrams[i*6+j]]!=row[hexagram[j]])
			if (row[hexagrams[i*6+j]]!=row[hexagram[j]]+2)
			if (row[hexagrams[i*6+j]]!=row[hexagram[j]]-2)
				match=FALSE;
		}
		if (match)
			index=i;
	}
	
	printf("%d\t%s\n\r", index+1, text$[index]);
	
	getchar();
	
	for (i=0;i<6;i++)
		hexagram2[i]=hexagram[i]%2;
	
	for (i=0;i<6;i++)
		printf("%s\n\r", row$[hexagram2[i]]);
		
	getchar();
				
	for (i=0;i<64;i++) {
		match=TRUE;
		for (j=0;j<6;j++) {
			if (row[hexagrams[i*6+j]]!=row[hexagram2[j]])
			if (row[hexagrams[i*6+j]]!=row[hexagram2[j]]+2)
			if (row[hexagrams[i*6+j]]!=row[hexagram2[j]]-2)
				match=FALSE;
		}
		if (match)
			index=i;
	}
	
	printf("%d\t%s\n\r", index+1, text$[index]);

	getchar();	
	return 0;
}
*/
