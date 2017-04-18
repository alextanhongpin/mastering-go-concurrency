package main

import (
	"fmt"
	"runtime"
	"strings"
)

var loremIpsum string
var finalIpsum string
var letterSentChannel chan string

func deliverToFinal(letter string, finalIpsum *string) {
	*finalIpsum += letter
}

func capitalize(current *int, length int, letters []byte, finalIpsum *string) {
	for *current < length {
		thisLetter := strings.ToUpper(string(letters[*current]))
		deliverToFinal(thisLetter, finalIpsum)
		*current++
	}
}

func main() {

	runtime.GOMAXPROCS(2)
	index := new(int)
	*index = 0

	loremIpsum = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut consequat, augue laoreet posuere pretium, velit ligula sagittis augue, non aliquam mauris sapien ut turpis. Praesent lacinia fringilla nulla, id ultrices dui auctor eu. Nulla nec ultrices ligula, a semper eros. Morbi consequat libero at convallis viverra. Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Fusce suscipit nibh ac purus vehicula, quis vehicula lectus posuere. Nunc imperdiet, neque eu viverra hendrerit, ante ante congue nibh, in luctus leo nulla in quam. Quisque quis augue erat. Curabitur placerat sagittis diam, non mattis neque lacinia ut. Sed sagittis consequat consectetur. Donec iaculis erat ac mauris pulvinar, sed molestie quam luctus. Nam tempus faucibus commodo. Nam lacinia est non justo placerat vestibulum in id massa. Integer consequat justo lacus, sit amet dignissim urna mattis eu. Cras fringilla metus porttitor eleifend pharetra. Mauris a aliquet metus, vel vestibulum elit. Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Suspendisse potenti. Aliquam vel posuere augue, a feugiat diam. Ut quis malesuada turpis, a gravida nunc. Suspendisse et metus ac mauris rutrum tincidunt in eget enim. Maecenas eu tortor vitae quam vestibulum efficitur. Nulla facilisi. Ut varius diam eget velit accumsan, non condimentum dui sollicitudin. Nunc ut diam porttitor mauris feugiat ornare. Vivamus et tellus vitae est finibus venenatis sed et ex. Sed dapibus semper augue, non rhoncus quam fringilla at. Vivamus massa quam, eleifend sed egestas quis, mattis in felis. Integer et nunc purus. Donec fringilla ullamcorper urna, luctus interdum tortor suscipit sed. Vivamus semper turpis et aliquam ullamcorper. In ut urna at nunc lacinia tempor. Nulla at vehicula sapien, in venenatis neque. Maecenas ultricies mollis erat id bibendum."

	letters := []byte(loremIpsum)
	length := len(letters)

	go capitalize(index, length, letters, &finalIpsum)
	go func() {
		capitalize(index, length, letters, &finalIpsum)
	}()

	fmt.Println(length, " characters.")
	fmt.Println(loremIpsum)
	fmt.Println(*index)
	fmt.Println(finalIpsum)
}
