
#include <stdio.h>
#include <stdlib.h>

#define SIZE 4

int* runningSum(int* nums, int numsSize, int* returnSize){
  
  // &[](addr) -> nums(addr) -> *[](value)
  // &[](addr) -> returnSize(addr) -> *[](value)
  // &[](addr) -> numsSize(value)
  // size of array
  // size of new array
  
  *returnSize = numsSize; 

  int* newNums = malloc(sizeof(int) * (*returnSize));
  newNums[0] = nums[0];

  // [1,2,3]
  // [1, 1+2, 1+2+3]
  
  for (int i =1; i < numsSize; i++){
    newNums[i] = newNums[i-1] + nums[i];
  }

  return newNums;

}

int main(){

  int nums[SIZE] = {1,2,3,4};
  int newSize = 1; // idk why it exist
  int* newNums = runningSum(nums, SIZE, &newSize);

  printf("[");
  for (int i =0; i < SIZE-1; i++){
    printf("%d,", newNums[i]);
  }
  printf("%d]\n", newNums[SIZE-1]);
  
  return 0;
}
