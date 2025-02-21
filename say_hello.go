package golangsayhello

func SayHello(name string) string {
	return "Hai" + name
}

// cmd git baru, setelah kita push ke repo, kita bs rilis tag versi
// git tag v1.0.0
// git push origin v1.0.0

// disarankan klo kita udah push, ganti versi tagnya walaupun minor, supaya perubahan kesimpan
// misalkan v1.0.1

// note jika ada major upgrade seperti perubahan pada isi program kita, sehingga membuatnya tidak backward compatible
// dimana haal ini sebisa mungkin dihindari
// namun jika g bs dihindari, strategi terbaik adalah merubah nama module
