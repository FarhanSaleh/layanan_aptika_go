package constants

const (
	// Sukses
	SuccessInsert  = "Data berhasil ditambahkan."
	SuccessUpdate  = "Data berhasil diperbarui."
	SuccessDelete  = "Data berhasil dihapus."
	SuccessGetData = "Data berhasil diambil."
	SuccessLogin   = "Berhasil masuk."
	SuccessLogout  = "Berhasil keluar."

	// Error Umum
	ErrorValidation     = "Data yang diberikan tidak valid."
	ErrorNotFound       = "Data tidak ditemukan."
	ErrorInsertFailed   = "Gagal menambahkan data."
	ErrorUpdateFailed   = "Gagal memperbarui data."
	ErrorDeleteFailed   = "Gagal menghapus data."
	ErrorInternalServer = "Terjadi kesalahan pada server."

	// Error Autentikasi & Authorization
	ErrorInvalidLogin  = "Email atau kata sandi salah."
	ErrorUnauthorized  = "Akses tidak diizinkan."
	ErrorAccountExists = "Akun dengan email tersebut sudah terdaftar."
)