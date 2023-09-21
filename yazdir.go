package main

import "fmt"

/*
func main() {
	var ogrenci string = "Akif"
	var ogrenci2 = "mert"
	ogrenci3 := "ammar"
	fmt.Println(ogrenci, ogrenci2, ogrenci3)

	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
*/
/*
func main() {
	var insan string
	fmt.Println(insan)
	insan = "Akif"
	fmt.Println(insan)
}
*/
/*
func main() {
	var a, b, c, d = 1, 2, 3, "merhaba"
	fmt.Println(a, b, c, d)
}
*/
/*
func main() {
	var (
		a int = 1
		b string = "merhaba"
		c bool = false
	)
	fmt.Println(a, b, c)
}
*/

/*
func main() {
	const degistirilemez_sabit string = "Beni değiştiremezsin!"
	fmt.Println(degistirilemez_sabit)
	degistirilemez_sabit = "Beni değiştirdin!" //cannot assign to degistirilemez_sabit (neither addressable nor a map index expression)
	fmt.Println(degistirilemez_sabit)
}
*/
/*
func main() {
	var a, b string = "merhaba", "selam"
	fmt.Print(a, b)
	var x, y = 10, 20
	fmt.Print(x, y)
}
*/

/*
func main() {
	var a string = "selam"
	var b int = 10

	fmt.Printf("a 'nın türü %T ve değeri %v\n", a, a)
	fmt.Printf("b 'nın türü %T ve değeri %v\n", b, b)
}
*/

/*
func main() {
	var i = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)
}
*/
/*
Verb 	Description
INTEGER
%b 		Base 2
%d 		Base 1
%+d 	Base 10 and always show sign
%o 		Base 8
%O 		Base 8, with leading 0o
%x 		Base 16, lowercase
%X 		Base 16, uppercase
%#x 	Base 16, with leading 0x
%4d 	Pad with spaces (width 4, right justified)
%-4d 	Pad with spaces (width 4, left justified)
%04d 	Pad with zeroes (width 4

STRING
%s 		Prints the value as plain string
%q 		Prints the value as a double-quoted string
%8s 	Prints the value as plain string (width 8, right justified)
%-8s 	Prints the value as plain string (width 8, left justified)
%x 	P	rints the value as hex dump of byte values
% x 	Prints the value as hex dump with spaces

BOOLEAN
%t 		Value of the boolean operator in true or false format (same as using %v)

FLOAT
%e 	Scientific notation with 'e' as exponent
%f 	Decimal point, no exponent
%.2f 	Default width, precision 2
%6.2f 	Width 6, precision 2
%g 	Exponent as needed, only necessary digits
*/

/*
func main() {
	var a1 bool = true
	var a2 = true
	var a3 bool //varsayılan bool değeri false'dur
	a4 := true

	fmt.Println(a1, a2, a3, a4)
}
*/

/*
func main() {
	dizi := [3]int{1, 5, 8}
	var arry = [2]string{"selam", "naber"}
	fmt.Println(dizi, arry)
}
*/

/*
func main() {
	dizi := [...]int{1, 5, 8}
	var arry = [...]string{"selam", "naber"}
	fmt.Println(dizi, arry)
}
*/
/*
func main() {
	dizi := [3]string{"ammar", "osman", "reyyan"}
	fmt.Println(dizi[2])
}
*/
/*
func main() {
	dizi := [3]string{"ammar", "osman", "reyyan"}
	fmt.Println(dizi)
	dizi[2] = "amede"
	fmt.Println(dizi)
}
*/
/*
func main() {
	dizi := [5]string{2: "ammar", 4: "osman", 1: "reyyan"}
	fmt.Println(dizi)
}
*/
/*
func main() {
	arr1 := [5]int{1: 10, 2: 40}

	fmt.Println(arr1)
}
*/
/*
func main() {
	dizi := [...]int{1, 5, 8}
	var arry = [...]string{"selam", "naber"}
	fmt.Println(len(dizi))
	fmt.Println(len(arry))
}
*/
/*
func main() {
	var arabalar = [4]string{"dacia", "reno", "mercedes", "togg"}
	fmt.Println(arabalar)
}
*/
/*
func main() {
	var arabalar = [4]string{"dacia", "reno", "mercedes", "togg"}
	fmt.Println(arabalar[2])
}
*/
/*
func main() {
	arabalar := [4]string{"dacia", "reno", "mercedes", "togg"}
	fmt.Println(arabalar)

	arabalar[2] = "tofaş"
	fmt.Println(arabalar)
}
*/
/*
func main() {
	arr1 := [5]int{}              //bomboş
	arr2 := [5]int{1, 2, 3}       //bir kısmı boş
	arr3 := [5]int{1, 2, 3, 4, 5} //dolu

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}
*/

//Initialize Only Specific Elements
/*
func main() {
	arr1 := [5]int{1: 20, 2: 20}

	fmt.Println(arr1)
}
*/
/*
func main() {
	arr1 := [...]string{"akif", "emre", "kayhan"}
	arr2 := [...]int{1, 23, 45, 44}

	fmt.Println(len(arr1), len(arr2))
}
*/

//SLICES
/*
func main() {
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
}
*/
/*
func main() {
	var arr1 = [3]int{1, 2, 5}
	slc1 := arr1[0:2]

	fmt.Printf("slice = %v\n", slc1)
	fmt.Printf("length = %d\n", len(slc1))
	fmt.Printf("capacity = %d\n", cap(slc1))
}
*/
/*
func main() {
	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	// with omitted capacity
	myslice2 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))
}
*/
/*
func main() {
	prices := []int{10, 20, 30}

	fmt.Println((prices[0]))
	fmt.Println((prices[2]))

}
*/
/*
func main() {
	prices := []int{10, 20, 30}

	fmt.Println((prices[0]))
	fmt.Println((prices[2]))

}
*/
/*
func main() {
	prices := []int{10, 20, 30}
	prices[2] = 50
	fmt.Println((prices[0]))
	fmt.Println((prices[2]))

}
*/
/*
func main() {
	prices := []int{10, 20, 30}
	fmt.Println(prices)
	prices = append(prices, 60, 70)
	fmt.Println(prices)

}
*/
/*
func main() {
	prices := []int{10, 20, 30}
	fmt.Println(prices)
	prices2 := []int{3, 5, 7}
	fmt.Println(prices2)
	prices3 := append(prices, prices2...)
	fmt.Println(prices3)
}
*/
/*
func main() {
	dizin := [8]int{5, 6, 7, 8, 9}
	fmt.Printf("Dizin içeriği %v\n", dizin)
	fmt.Printf("Dizin uzunluğu %d\n", len(dizin))
	fmt.Printf("Dizin kapasitesi %d\n", cap(dizin))

	dilim := dizin[:4]
	fmt.Printf("Dilim içeriği %v\n", dilim)
	fmt.Printf("Dilim uzunluğu %d\n", len(dilim))
	fmt.Printf("Dilim kapasitesi %d\n", cap(dilim))
}
*/

//BELLEK YÖNETİMİ
/*
func main() {
	rakamlar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Printf("numbers = %v\n", rakamlar)
	fmt.Printf("length = %d\n", len(rakamlar))
	fmt.Printf("capacity = %d\n", cap(rakamlar))

	//Belleği yormamak içins sadece gerekli öğeleri al
	alınanOgeler := rakamlar[:len(rakamlar)-10]
	kopyalananOgeler := make([]int, len(alınanOgeler))
	copy(kopyalananOgeler, alınanOgeler)

	fmt.Printf("kopyalananOgeler = %v\n", kopyalananOgeler)
	fmt.Printf("uzunluk = %d\n", len(alınanOgeler))
	fmt.Printf("kapasite = %d\n", cap(kopyalananOgeler))
}
*/

//OPERATÖRLER
/*
func main() {
	a := 15 + 10
	fmt.Println(a)
}
*/
/*
func main() {
	var (
		toplam1 = 35 + 35
		toplam2 = toplam1 + 40
		toplam3 = toplam2 + toplam1
	)
	fmt.Println(toplam3)
}
*/

/*
Operator 	Name 			Description 							Example
+ 			Addition 		Adds together two values 				x + y
- 			Subtraction 	Subtracts one value from another 		x - y
* 			Multiplication 	Multiplies two values 					x * y
/ 			Division 		Divides one value by another 			x / y
% 			Modulus 		Returns the division remainder 			x % y
++ 			Increment 		Increases the value of a variable by 1 	x++
-- 			Decrement 		Decreases the value of a variable by 1 	x--
*/

//GÖREVLENDİRME OPERATÖRLERİ
/*
func main() {
	//Atama operatörleri değişkenlere değer atamak için kullanılır.
	var x = 10
	fmt.Println(x)
  }
*/
/*
func main() {
	//Toplama atama operatörü (+=) bir değişkene değer ekler:
	var x = 10
	x +=5
	fmt.Println(x)
}
*/
/*
Operator 	Example 	Same As
= 			x = 5 		x = 5
+= 			x += 3 		x = x + 3
-= 			x -= 3 		x = x - 3
*= 			x *= 3 		x = x * 3
/= 			x /= 3 		x = x / 3
%= 			x %= 3 		x = x % 3
&= 			x &= 3 		x = x & 3
|= 			x |= 3 		x = x | 3
^= 			x ^= 3 		x = x ^ 3
>>= 		x >>= 3 	x = x >> 3
<<= 		x <<= 3 	x = x << 3

&= (Bitwise AND ve Atama):
x &= y, x ve y'nin her bir bitini AND işlemine tabi tutar ve sonucu x'e atar.
Örnek: x &= 3, x'in bütün bitlerini 3 ile AND işlemine tabi tutar ve sonucu x'e atar.

|= (Bitwise OR ve Atama):
x |= y, x ve y'nin her bir bitini OR işlemine tabi tutar ve sonucu x'e atar.
Örnek: x |= 3, x'in bütün bitlerini 3 ile OR işlemine tabi tutar ve sonucu x'e atar.

^= (Bitwise XOR ve Atama):
x ^= y, x ve y'nin her bir bitini XOR işlemine tabi tutar ve sonucu x'e atar.
Örnek: x ^= 3, x'in bütün bitlerini 3 ile XOR işlemine tabi tutar ve sonucu x'e atar.

>>= (Sağa Kaydırma ve Atama):
x >>= y, x'in bütün bitlerini sağa y adım kaydırır ve sonucu x'e atar.
Örnek: x >>= 3, x'in bütün bitlerini 3 adım sağa kaydırır ve sonucu x'e atar.

<<= (Sola Kaydırma ve Atama):
x <<= y, x'in bütün bitlerini sola y adım kaydırır ve sonucu x'e atar.
Örnek: x <<= 3, x'in bütün bitlerini 3 adım sola kaydırır ve sonucu x'e atar.
*/

//KARŞILAŞTIRMA OPERATÖRLERİ

/*
func main() {
	var x = 5
	var y = 3
	fmt.Println(x>y) // returns 1 (true) because 5 is greater than 3
}
*/

/*
Operator 	Name 						Example
== 			Equal to 					x == y
!= 			Not equal 					x != y
> 			Greater than 				x > y
< 			Less than 					x < y
>= 			Greater than or equal to 	x >= y
<= 			Less than or equal to 		x <= y
*/

//MANTIKSAL OPERATÖRLER
/*
Mantıksal operatörler, değişkenler veya değerler arasındaki mantığı belirlemek için kullanılır:

Operator 		Name 			Description 												Example
&&  			Logical and 	Returns true if both statements are true 					x < 5 &&  x < 10
||  			Logical or 		Returns true if one of the statements is true 				x < 5 || x < 4
! 				Logical not 	Reverse the result, returns false if the result is true 	!(x < 5 && x < 10)
*/

/*
BİTSEL OPERATÖRLER
Bitsel operatörler (ikili) sayılar üzerinde kullanılır:

Operator	Name 					Description 										Example
&  			AND 					Sets each bit to 1 if both bits are 1 				x & y
| 			OR 						Sets each bit to 1 if one of two bits is 1 			x | y
^ 			XOR 					Sets each bit to 1 if only one of two bits is 1 	x ^ b
<< 			Zero fill left shift 	Shift left by pushing zeros in from the right 		x << 2
>> 			Signed right shift 		Shift right by pushing copies of the leftmost		x >> 2
									bit in from the left, and let the rightmost
									bits fall off
*/

//DURUMLAR
/*
func main() {
	temperature := 14
	if (temperature > 15) {
	  fmt.Println("It is warm out there")
	} else {
	  fmt.Println("It is cold out there")
	}
  }
*/

//SWITCH-CASE

/*
func main() {
	day := 4

	switch day  {
	case 1:
	  fmt.Println("Monday")
	case 2:
	  fmt.Println("Tuesday")
	case 3:
	  fmt.Println("Wednesday")
	case 4:
	  fmt.Println("Thursday")
	case 5:
	  fmt.Println("Friday")
	case 6:
	  fmt.Println("Saturday")
	case 7:
	  fmt.Println("Sunday")
	default:
	  fmt.Println("Not a weekday")
  }
}
*/
/*
  func main() {
	day := 5

	switch day {
	case 1,3,5:
	 fmt.Println("Odd weekday")
	case 2,4:
	  fmt.Println("Even weekday")
	case 6,7:
	 fmt.Println("Weekend")
   default:
	 fmt.Println("Invalid day of day number")
   }
 }*/

//FOR DÖNGÜLERİ
/*
func main() {
	for i := 5/*başlangıç değeri*/ //; i < 255/*koşul*/; i *= 5/*çıkarım*/ /*{
//	fmt.Println(i)
/*
	}
}
*/
/*
func main() {
	for i := 2; i < 1025; i *= 3 {
		if 250 < i && i < 500 {
			continue
		}
		fmt.Println(i)
	}
}
*/

/*func main() {
	for i := 2; i < 1025; i *= 3 {
		if 250 < i && i < 500 {
			break
		}
		fmt.Println(i)
	}
}*/

//nested loops
/*
func main() {
	renk := [3]string{"mavi", "kırmızı", "sarı"}
	arac := [4]string{"araba", "bisiklet", "uçak", "motosiklet"}
	for i := 0; i < len(renk); i++ {
		for j := 0; j < len(arac); j++ {
			fmt.Println(renk[i], arac[j])
		}
	}
}
*/

//ARALIK
/*
func main() {
	arac := [4]string{"araba", "bisiklet", "uçak", "motosiklet"}
	for sira, deger := range arac {
		fmt.Printf("%v\t%v\n", sira, deger)
	}
}
*/
/*
func main() {
	arac := [4]string{"araba", "bisiklet", "uçak", "motosiklet"}
	for _, deger := range arac {
		fmt.Printf("%v\n", deger)
	}
}
*/
/*
func main() {
	arac := [4]string{"araba", "bisiklet", "uçak", "motosiklet"}
	for sira := range arac {
		fmt.Printf("%v\n", sira)
	}
}
*/

// FONKSİYONLAR
/*
func fonksiyon() {
	fmt.Println("Fonksiyon çağırıldı!")
}

func main() {
	fonksiyon()
}
*/
/*
func aile(aisim string) {
	fmt.Println("selam", aisim, "kayhan")
}

func main() {
	aile("akif emre")
	aile("semih")
	aile("sema")
	aile("zülal")
	aile("dilara")
}
*/
/*
func hangitakimli(takim string, isim string) {
	fmt.Println(isim, takim, "lı")
}

func main() {
	hangitakimli("beşiktaş", "emir")
	hangitakimli("galatasaray", "yahya")
	hangitakimli("fenerbahçe", "hamza")
}
*/
/*
func toplama(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(toplama(5, 6))
}
*/
/*
func toplama(x int, y int) (toplam int) {
	toplam = x + y
	return
}

func main() {
	fmt.Println(toplama(5, 6))
}
*/
/*
func toplama(x int, y int) (toplam int) {
	toplam = x + y
	return
}

func main() {
	sonuc := toplama(5, 6)
	fmt.Println(sonuc)
}
*/
/*
func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func main() {
	fmt.Println(myFunction(5, "Hello"))
}
*/
/*
func artar(x int) int {
	if x == 256 {
		return 0
	}
	fmt.Println(x)
	return artar(x * 2)
}

func main() {
	artar(1)
}
*/
/*
func faktoriyel(x int) (y int) {
	if x != 0 {
		y = x*faktoriyel(x-1)
	} else {
		y = 1
	}
	return
}

func main() {
	fmt.Println(faktoriyel(6))
}
*/
/*
func factorial(n int) int {
	// Temel Durum: n 0 veya 1 ise, faktöriyel 1'dir.
	if n <= 1 {
		return 1
	}
	// Rekürsif Adım: n faktöriyel, n ile (n-1) faktöriyelinin çarpımıdır.
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(4))
}
*/

//STRUCT
/*
type Ogrenci struct {
	isim     string
	yas      int
	okul     string
	ortalama float64
}

func main() {
	var akif Ogrenci
	var menes Ogrenci

	akif.isim = "Akif Emre"
	akif.yas = 15
	akif.okul = "kaihl"
	akif.ortalama = 90.4

	menes.isim = "Muhammet Enes"
	menes.yas = 16
	menes.okul = "ikhal"
	menes.ortalama = 90.7

	ogrenciYazdir(akif)
	ogrenciYazdir(menes)
}

func ogrenciYazdir(kisi Ogrenci) {
	fmt.Println("İsim: ", kisi.isim, "\nYaş: ", kisi.yas, "\nOkul: ", kisi.okul, "\nOrtalama: ", kisi.ortalama)
}
*/

//MAP'ler
/*
func main() {
	harita := map[string]int{"akif": 1, "ammar": 2, "apo": 3}
	fmt.Printf("harita\t%v\n", harita)
}
*/
/*
func main() {
	var harita = make(map[string]int)
	harita["akif"] = 1
	harita["ammar"] = 2
	harita["apo"] = 3

	fmt.Printf("harita\t%v\n", harita)
}
*/
/*
func main() {
	var harita map [string]int
*/
/*
func main() {
	var harita = make(map[string]int)
	harita["akif"] = 1
	harita["ammar"] = 2
	harita["apo"] = 3

	fmt.Println(harita["akif"])
}
*/
/*
func main() {
	var harita = make(map[string]int)
	harita["akif"] = 1
	harita["ammar"] = 2
	harita["apo"] = 3

	harita["yavuz"] = 4
	harita["ammar"] = 5
	fmt.Println(harita)
}
*/
/*
func main() {
	var harita = make(map[string]int)
	harita["akif"] = 1
	harita["ammar"] = 2
	fmt.Println(harita)

	delete(harita, "akif")
	fmt.Println(harita)
}
*/
/*
func main() {
	var harita = map[string]string{"akif": "kayhan", "ammar": "", "apo": "güveren"}

	deger1, ok1 := harita["akif"]
	deger2, ok2 := harita["ammar"]
	_, ok3 := harita["yokbolebisi"]
	fmt.Println(deger1, ok1)
	fmt.Println(deger2, ok2)
	fmt.Println(ok3)
}
*/
/*
func main() {
	var harita = map[string]string{"akif": "kayhan", "ammar": "", "apo": "güveren"}
	haritamsi := harita

	fmt.Println(haritamsi)
}
*/
/*
func main() {
	sira_sayilari := map[string]int{"bir": 1, "iki": 2, "üç": 3, "dört": 4}

	for metin, sayi := range sira_sayilari {
		fmt.Printf("%v : %v\n", metin, sayi)
	}
}
*/

func main() {
	sira_sayilari := map[string]int{"bir": 1, "iki": 2, "üç": 3, "dört": 4}

	var secme_sayilar []string
	secme_sayilar = append(secme_sayilar, "bir", "üç", "dört")

	for metin, sayi := range sira_sayilari {
		fmt.Printf("%v : %v\n", metin, sayi)
	}

	for _, element := range secme_sayilar {
		fmt.Printf("%v : %v\n", element, sira_sayilari[element])
	}
}
