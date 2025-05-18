package main

import "fmt"

type UserProfile struct {
	Nama            string
	Email           string
	Phone           string
	Alamat          string
	PengalamanKerja [10]Experience
	Skills          [20]Skill
	Education       [5]Education
	WorkExpCount    int
	SkillCount      int
	EduCount        int
}

type Experience struct {
	Perusahaan string
	Position   string
	StartYear  int
	EndYear    int
	Deskripsi  string
}

type Skill struct {
	Nama      string
	Level     int
	Deskripsi string
}

type Education struct {
	Institusi  string
	Degree     string
	Bidang     string
	TahunLulus int
}

type Job struct {
	Title        string
	Perusahaan   string
	Deskripsi    string
	Keywords     [10]string
	KeywordCount int
	Salary       int
	Relevance    int
}

type tabJob [20]Job

func main() {
	var profile UserProfile
	var jobs tabJob
	var jobCount int
	var pilihan int = -1

	jobs[0] = Job{
		Title:        "Software Engineer",
		Perusahaan:   "PT. Gobot Teknologi",
		Deskripsi:    "Mengembangkan aplikasi web menggunakan Golang dan React",
		Keywords:     [10]string{"golang", "react", "git", "agile", "backend", "frontend", "", "", "", ""},
		KeywordCount: 6,
		Salary:       15000000,
		Relevance:    0,
	}

	jobs[1] = Job{
		Title:        "Data Analyst",
		Perusahaan:   "PT. Atmapurwa Global Teknologi",
		Deskripsi:    "Menganalisis data perusahaan dan membuat visualisasi",
		Keywords:     [10]string{"sql", "excel", "python", "statistics", "visualization", "", "", "", "", ""},
		KeywordCount: 5,
		Salary:       12000000,
		Relevance:    0,
	}

	jobs[2] = Job{
		Title:        "Product Manager",
		Perusahaan:   "Inovasi Digital",
		Deskripsi:    "Mengelola roadmap produk dan berkoordinasi dengan tim teknis",
		Keywords:     [10]string{"product", "management", "agile", "roadmap", "communication", "leadership", "", "", "", ""},
		KeywordCount: 6,
		Salary:       20000000,
		Relevance:    0,
	}
	jobCount = 3

	for pilihan != 0 {
		tampilkanMenuUtama(&pilihan)

		if pilihan == 1 {
			kelolaProfil(&profile)
		} else if pilihan == 2 {
			kelolaCV(&profile, jobs, jobCount)
		} else if pilihan == 3 {
			cariPekerjaan(jobs, jobCount)
		} else if pilihan == 4 {
			urutkanPekerjaan(jobs, jobCount)
		} else if pilihan == 5 {
			evaluasiResume(&profile, jobs, jobCount)
		} else if pilihan == 0 {
			fmt.Println("Terima kasih telah menggunakan aplikasi kami")
		}
	}

}

func tampilkanMenuUtama(pilihan *int) {
	fmt.Print("\n*** APLIKASI PEMBUAT RESUME DAN SURAT LAMARAN ***\nHaloo apakah ada yang bisa DafNat AI bantu? aku punya beberapa opsi pilihan untuk kamu\n",
		" \n1. Kelola Profil",
		" \n2. Buat Resume dan Surat Lamaran",
		" \n3. Cari Pekerjaan",
		" \n4. Urutkan Daftar Pekerjaan",
		" \n5. Evaluasi Resume",
		" \n0. Keluar",
		" \nPilihan Anda: ",
	)
	fmt.Scan(pilihan)
}

func kelolaProfil(profile *UserProfile) {
	var isManagingProfile bool = true
	for isManagingProfile {

		fmt.Print("\n*** KELOLA PROFIL ***",
			" \n1. Lihat Profil",
			" \n2. Ubah Data Diri",
			" \n3. Kelola Pengalaman Kerja",
			" \n4. Kelola Keterampilan",
			" \n5. Kelola Pendidikan",
			" \n0. Kembali ke Menu Utama",
			" \nPilihan Anda: ",
		)

		var pilihan int = -1
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			lihatProfil(*profile)
		} else if pilihan == 2 {
			ubahDataDiri(profile)
		} else if pilihan == 3 {
			kelolaPengalamanKerja(profile)
		} else if pilihan == 4 {
			kelolaKeterampilan(profile)
		} else if pilihan == 5 {
			kelolaPendidikan(profile)
		} else if pilihan == 0 {
			isManagingProfile = false
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}

	}
}

func lihatProfil(profile UserProfile) {
	var i int

	fmt.Println("\n*** PROFIL ANDA ***")
	fmt.Println("Nama     :", profile.Nama)
	fmt.Println("Email    :", profile.Email)
	fmt.Println("Telepon  :", profile.Phone)
	fmt.Println("Alamat   :", profile.Alamat)

	fmt.Println("\n--- PENGALAMAN KERJA ---")
	if profile.WorkExpCount == 0 {
		fmt.Println("Belum ada pengalaman kerja.")
	} else {
		for i = 0; i < profile.WorkExpCount; i++ {
			fmt.Printf("%d. %s di %s (%d-%d)\n", i+1, profile.PengalamanKerja[i].Position, profile.PengalamanKerja[i].Perusahaan, profile.PengalamanKerja[i].StartYear, profile.PengalamanKerja[i].EndYear)
			fmt.Printf("   Deskripsi: %s\n", profile.PengalamanKerja[i].Deskripsi)
		}
	}

	fmt.Println("\n--- KETERAMPILAN ---")
	if profile.SkillCount == 0 {
		fmt.Println("Belum ada keterampilan.")
	} else {
		for i = 0; i < profile.SkillCount; i++ {
			fmt.Printf("%d. %s (Level %d)\n", i+1, profile.Skills[i].Nama, profile.Skills[i].Level)
			fmt.Printf("   Deskripsi: %s\n", profile.Skills[i].Deskripsi)
		}
	}

	fmt.Println("\n--- PENDIDIKAN ---")
	if profile.EduCount == 0 {
		fmt.Println("Belum ada pendidikan.")
	} else {
		for i = 0; i < profile.EduCount; i++ {
			fmt.Printf("%d. %s di %s, %s (%d)\n", i+1, profile.Education[i].Degree, profile.Education[i].Institusi, profile.Education[i].Bidang, profile.Education[i].TahunLulus)
		}
	}

	fmt.Println("\nSaya siap saat Anda siap. Tekan Enter untuk melanjutkan...")
	var input string
	fmt.Scanln(&input)
}

func ubahDataDiri(profile *UserProfile) {
	fmt.Println("\n*** UBAH DATA DIRI ***")

	var data string

	fmt.Printf("Nama (%s): ", profile.Nama)
	fmt.Scanln(&data)
	if data != "" {
		profile.Nama = data
	}

	fmt.Printf("Email (%s): ", profile.Email)
	data = ""
	fmt.Scanln(&data)
	if data != "" {
		profile.Email = data
	}

	fmt.Printf("Telepon (%s): ", profile.Phone)
	data = ""
	fmt.Scanln(&data)
	if data != "" {
		profile.Phone = data
	}

	fmt.Printf("Alamat (%s): ", profile.Alamat)
	data = ""
	fmt.Scanln(&data)
	if data != "" {
		profile.Alamat = data
	}

	fmt.Println("Data diri berhasil diperbarui!")
}

func kelolaPengalamanKerja(profile *UserProfile) {
	var isManagingExp bool = true
	for isManagingExp {
		fmt.Print("\n*** KELOLA PENGALAMAN KERJA ***",
			" \n1. Lihat Pengalaman Kerja",
			" \n2. Tambah Pengalaman Kerja",
			" \n3. Ubah Pengalaman Kerja",
			" \n4. Hapus Pengalaman Kerja",
			" \n0. Kembali",
			" \nPilihan Anda: ",
		)

		var pilihan int = -1
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			lihatPengalamanKerja(*profile)
		} else if pilihan == 2 {
			tambahPengalamanKerja(profile)
		} else if pilihan == 3 {
			ubahPengalamanKerja(profile)
		} else if pilihan == 4 {
			hapusPengalamanKerja(profile)
		} else if pilihan == 5 {
			kelolaPendidikan(profile)
		} else if pilihan == 0 {
			isManagingExp = false
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}

	}
}

func lihatPengalamanKerja(profile UserProfile) {
	var i int
	fmt.Println("\n*** PENGALAMAN KERJA ANDA ***")
	if profile.WorkExpCount == 0 {
		fmt.Println("Belum ada pengalaman kerja yang ditambahkan.")
		return
	}

	for i = 0; i < profile.WorkExpCount; i++ {
		fmt.Printf("%d. %s di %s (%d-%d)\n", i+1, profile.PengalamanKerja[i].Position, profile.PengalamanKerja[i].Perusahaan, profile.PengalamanKerja[i].StartYear, profile.PengalamanKerja[i].EndYear)
		fmt.Printf("   Deskripsi: %s\n", profile.PengalamanKerja[i].Deskripsi)
	}

	fmt.Println("\nSaya siap saat Anda siap. Tekan Enter untuk melanjutkan...")
	var input string
	fmt.Scanln(&input)
}

func tambahPengalamanKerja(profile *UserProfile) {
	if profile.WorkExpCount >= 10 {
		fmt.Println("Batas maksimum pengalaman kerja tercapai (10).")
		return
	}

	fmt.Println("\n*** TAMBAH PENGALAMAN KERJA ***")
	fmt.Print("Nama Perusahaan: ")
	fmt.Scan(&profile.PengalamanKerja[profile.WorkExpCount].Perusahaan)

	fmt.Print("Posisi/Jabatan: ")
	fmt.Scan(&profile.PengalamanKerja[profile.WorkExpCount].Position)

	fmt.Print("Tahun Mulai: ")
	fmt.Scan(&profile.PengalamanKerja[profile.WorkExpCount].StartYear)

	fmt.Print("Tahun Berakhir (masukkan 0 jika masih bekerja di sini): ")
	fmt.Scan(&profile.PengalamanKerja[profile.WorkExpCount].EndYear)

	fmt.Print("Deskripsi Pekerjaan: ")
	fmt.Scan(&profile.PengalamanKerja[profile.WorkExpCount].Deskripsi)

	profile.WorkExpCount += 1

	fmt.Println("Pengalaman kerja berhasil ditambahkan!")
}

func ubahPengalamanKerja(profile *UserProfile) {
	if profile.WorkExpCount == 0 {
		fmt.Println("Belum ada pengalaman kerja yang ditambahkan.")
		return
	}

	lihatPengalamanKerja(*profile)

	fmt.Print("\nPilih nomor pengalaman kerja yang ingin diubah: ")
	var index int
	fmt.Scan(&index)

	if index < 1 || index > profile.WorkExpCount {
		fmt.Println("Nomor tidak valid.")
		return
	}

	index -= 1

	fmt.Printf("Nama Perusahaan (%s): ", profile.PengalamanKerja[index].Perusahaan)
	fmt.Scan(&profile.PengalamanKerja[index].Perusahaan)

	fmt.Printf("Posisi/Jabatan (%s): ", profile.PengalamanKerja[index].Position)
	fmt.Scan(&profile.PengalamanKerja[index].Position)

	fmt.Printf("Tahun Mulai (%d): ", profile.PengalamanKerja[index].StartYear)
	fmt.Scan(&profile.PengalamanKerja[index].StartYear)

	fmt.Printf("Tahun Berakhir (%d): ", profile.PengalamanKerja[index].EndYear)
	fmt.Scan(&profile.PengalamanKerja[index].EndYear)

	fmt.Printf("Deskripsi Pekerjaan (%s): ", profile.PengalamanKerja[index].Deskripsi)
	fmt.Scan(&profile.PengalamanKerja[index].Deskripsi)

	fmt.Println("Pengalaman kerja berhasil diperbarui!")
}

func hapusPengalamanKerja(profile *UserProfile) {
	var i int

	if profile.WorkExpCount == 0 {
		fmt.Println("Belum ada pengalaman kerja yang ditambahkan.")
		return
	}

	lihatPengalamanKerja(*profile)

	fmt.Print("\nPilih nomor pengalaman kerja yang ingin dihapus: ")
	var index int
	fmt.Scan(&index)

	if index < 1 || index > profile.WorkExpCount {
		fmt.Println("Nomor tidak valid.")
		return
	}

	index -= 1

	for i = index; i < profile.WorkExpCount-1; i++ {
		profile.PengalamanKerja[i] = profile.PengalamanKerja[i+1]
	}

	profile.WorkExpCount -= 1
	fmt.Println("Pengalaman kerja berhasil dihapus!")
}

func kelolaKeterampilan(profile *UserProfile) {
	var isManagingSkills bool = true
	for isManagingSkills {

		fmt.Print("\n*** KELOLA KETERAMPILAN ***",
			" \n1. Lihat Keterampilan",
			" \n2. Tambah Keterampilan",
			" \n3. Ubah Keterampilan",
			" \n4. Hapus Keterampilan",
			" \n0. Kembali",
			" \nPilihan Anda: ",
		)

		var pilihan int = -1
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			lihatKeterampilan(*profile)
		} else if pilihan == 2 {
			tambahKeterampilan(profile)
		} else if pilihan == 3 {
			ubahKeterampilan(profile)
		} else if pilihan == 4 {
			hapusKeterampilan(profile)
		} else if pilihan == 0 {
			isManagingSkills = false
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func lihatKeterampilan(profile UserProfile) {
	var i int
	fmt.Println("\n*** KETERAMPILAN ANDA ***")
	if profile.SkillCount == 0 {
		fmt.Println("Belum ada keterampilan yang ditambahkan.")
		return
	}

	for i = 0; i < profile.SkillCount; i++ {
		fmt.Printf("%d. %s (Level %d)\n", i+1, profile.Skills[i].Nama, profile.Skills[i].Level)
		fmt.Printf("   Deskripsi: %s\n", profile.Skills[i].Deskripsi)
	}

	fmt.Println("\nSaya siap saat Anda siap. Tekan Enter untuk melanjutkan...")
	var input string
	fmt.Scanln(&input)
}

func tambahKeterampilan(profile *UserProfile) {
	if profile.SkillCount >= 20 {
		fmt.Println("Batas maksimum keterampilan tercapai (20).")
		return
	}

	fmt.Println("\n*** TAMBAH KETERAMPILAN ***")
	fmt.Print("Nama Keterampilan: ")
	fmt.Scan(&profile.Skills[profile.SkillCount].Nama)

	fmt.Print("Level (1-5): ")
	fmt.Scan(&profile.Skills[profile.SkillCount].Level)

	if profile.Skills[profile.SkillCount].Level < 1 || profile.Skills[profile.SkillCount].Level > 5 {
		fmt.Println("Level harus antara 1-5. Defaultnya diatur ke 1.")
		profile.Skills[profile.SkillCount].Level = 1
	}

	fmt.Print("Deskripsi Keterampilan: ")
	fmt.Scan(&profile.Skills[profile.SkillCount].Deskripsi)

	profile.SkillCount += 1

	fmt.Println("Keterampilan berhasil ditambahkan!")
}

func ubahKeterampilan(profile *UserProfile) {
	if profile.SkillCount == 0 {
		fmt.Println("Belum ada keterampilan yang ditambahkan.")
		return
	}

	lihatKeterampilan(*profile)

	fmt.Print("\nPilih nomor keterampilan yang ingin diubah: ")
	var index int
	fmt.Scan(&index)

	if index < 1 || index > profile.SkillCount {
		fmt.Println("Nomor tidak valid.")
		return
	}

	index -= 1

	fmt.Printf("Nama Keterampilan (%s): ", profile.Skills[index].Nama)
	fmt.Scan(&profile.Skills[index].Nama)

	fmt.Printf("Level (%d): ", profile.Skills[index].Level)
	fmt.Scan(&profile.Skills[index].Level)
	if profile.Skills[index].Level != 0 {
		if profile.Skills[index].Level < 1 || profile.Skills[index].Level > 5 {
			fmt.Println("Level harus antara 1-5. Tidak diubah.")
		}
	}

	fmt.Printf("Deskripsi Keterampilan (%s): ", profile.Skills[index].Deskripsi)
	fmt.Scan(&profile.Skills[index].Deskripsi)

	fmt.Println("Keterampilan berhasil diperbarui!")
}

func hapusKeterampilan(profile *UserProfile) {
	var i int

	if profile.SkillCount == 0 {
		fmt.Println("Belum ada keterampilan yang ditambahkan.")
		return
	}

	lihatKeterampilan(*profile)

	fmt.Print("\nPilih nomor keterampilan yang ingin dihapus: ")
	var index int
	fmt.Scan(&index)

	if index < 1 || index > profile.SkillCount {
		fmt.Println("Nomor tidak valid.")
		return
	}

	index -= 1

	for i = index; i < profile.SkillCount-1; i++ {
		profile.Skills[i] = profile.Skills[i+1]
	}

	profile.SkillCount -= 1
	fmt.Println("Keterampilan berhasil dihapus!")
}

func kelolaPendidikan(profile *UserProfile) {
	var isManagingEducation bool = true
	for isManagingEducation {

		fmt.Print("\n*** KELOLA PENDIDIKAN ***",
			" \n1. Lihat Pendidikan",
			" \n2. Tambah Pendidikan",
			" \n3. Ubah Pendidikan",
			" \n4. Hapus Pendidikan",
			" \n0. Kembali",
			" \nPilihan Anda: ",
		)

		var pilihan int = -1
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			lihatPendidikan(*profile)
		} else if pilihan == 2 {
			tambahPendidikan(profile)
		} else if pilihan == 3 {
			ubahPendidikan(profile)
		} else if pilihan == 4 {
			hapusPendidikan(profile)
		} else if pilihan == 0 {
			isManagingEducation = false
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}

	}
}

func lihatPendidikan(profile UserProfile) {
	var i int
	fmt.Println("\n*** PENDIDIKAN ANDA ***")
	if profile.EduCount == 0 {
		fmt.Println("Belum ada pendidikan yang ditambahkan.")
		return
	}

	for i = 0; i < profile.EduCount; i++ {
		fmt.Printf("%d. %s di %s, %s (%d)\n", i+1, profile.Education[i].Degree, profile.Education[i].Institusi, profile.Education[i].Bidang, profile.Education[i].TahunLulus)
	}

	fmt.Println("\nSaya siap saat Anda siap. Tekan Enter untuk melanjutkan...")
	var input string
	fmt.Scanln(&input)
}

func tambahPendidikan(profile *UserProfile) {
	if profile.EduCount >= 5 {
		fmt.Println("Batas maksimum pendidikan tercapai (5).")
		return
	}

	fmt.Println("\n*** TAMBAH PENDIDIKAN ***")
	fmt.Print("Nama Institusi: ")
	fmt.Scan(&profile.Education[profile.EduCount].Institusi)

	fmt.Print("Gelar (contoh: S1, S2, Diploma): ")
	fmt.Scan(&profile.Education[profile.EduCount].Degree)

	fmt.Print("Bidang/Jurusan: ")
	fmt.Scan(&profile.Education[profile.EduCount].Bidang)

	fmt.Print("Tahun Lulus: ")
	fmt.Scan(&profile.Education[profile.EduCount].TahunLulus)

	profile.EduCount += 1

	fmt.Println("Pendidikan berhasil ditambahkan!")
}

func ubahPendidikan(profile *UserProfile) {
	if profile.EduCount == 0 {
		fmt.Println("Belum ada pendidikan yang ditambahkan.")
		return
	}

	lihatPendidikan(*profile)

	fmt.Print("\nPilih nomor pendidikan yang ingin diubah: ")
	var index int
	fmt.Scan(&index)

	if index < 1 || index > profile.EduCount {
		fmt.Println("Nomor tidak valid.")
		return
	}

	index -= 1

	fmt.Printf("Nama Institusi (%s): ", profile.Education[index].Institusi)
	fmt.Scan(&profile.Education[index].Institusi)

	fmt.Printf("Gelar (%s): ", profile.Education[index].Degree)
	fmt.Scan(&profile.Education[index].Degree)

	fmt.Printf("Bidang/Jurusan (%s): ", profile.Education[index].Bidang)
	fmt.Scan(&profile.Education[index].Bidang)

	fmt.Printf("Tahun Lulus (%d): ", profile.Education[index].TahunLulus)
	fmt.Scan(&profile.Education[index].TahunLulus)

	fmt.Println("Pendidikan berhasil diperbarui!")
}

func hapusPendidikan(profile *UserProfile) {
	var i int

	if profile.EduCount == 0 {
		fmt.Println("Belum ada pendidikan yang ditambahkan.")
		return
	}

	lihatPendidikan(*profile)

	fmt.Print("\nPilih nomor pendidikan yang ingin dihapus: ")
	var index int
	fmt.Scan(&index)

	if index < 1 || index > profile.EduCount {
		fmt.Println("Nomor tidak valid.")
		return
	}

	index -= 1

	for i = index; i < profile.EduCount-1; i++ {
		profile.Education[i] = profile.Education[i+1]
	}

	profile.EduCount -= 1
	fmt.Println("Pendidikan berhasil dihapus!")
}

func mengandungString(string, substring string) bool {
	var strLen, substringLen, i, j int
	var matched bool

	strLen = len(string)
	substringLen = len(substring)

	if substringLen > strLen {
		return false
	}

	for i = 0; i <= strLen-substringLen; i++ {
		matched = true
		for j = 0; j < substringLen; j++ {
			if string[i+j] != substring[j] {
				matched = false
			}
		}
		if matched {
			return true
		}
	}

	return false
}

func hitungRelevansiPekerjaan(profile *UserProfile, job *Job) {
	var i, j int

	var totalRelevance int = 0
	var matchedKeywords int = 0

	for i = 0; i < profile.WorkExpCount; i++ {

		for j = 0; j < job.KeywordCount; j++ {

			if mengandungString(profile.PengalamanKerja[i].Position, job.Keywords[j]) || mengandungString(profile.PengalamanKerja[i].Deskripsi, job.Keywords[j]) {
				matchedKeywords++
			}
		}
	}

	for i = 0; i < profile.SkillCount; i++ {

		for j = 0; j < job.KeywordCount; j++ {

			if mengandungString(profile.Skills[i].Nama, job.Keywords[j]) || mengandungString(profile.Skills[i].Deskripsi, job.Keywords[j]) {
				matchedKeywords++
				totalRelevance += profile.Skills[i].Level * 5
			}
		}
	}

	if job.KeywordCount > 0 {
		totalRelevance = (totalRelevance + (matchedKeywords * 10)) / (job.KeywordCount * 2)
	}

	if totalRelevance > 100 {
		totalRelevance = 100
	}

	job.Relevance = totalRelevance
}

func kelolaCV(profile *UserProfile, jobs tabJob, jobCount int) {
	var i int

	if profile.Nama == "" || profile.Email == "" || profile.Phone == "" {
		fmt.Println("\nAnda perlu mengisi informasi profil dasar terlebih dahulu.")
		return
	}

	if profile.WorkExpCount == 0 && profile.SkillCount == 0 && profile.EduCount == 0 {
		fmt.Println("\nAnda perlu menambahkan minimal satu pengalaman kerja, keterampilan, atau pendidikan.")
		return
	}

	fmt.Println("\n*** BUAT RESUME DAN SURAT LAMARAN ***")
	fmt.Println("Pilih pekerjaan yang ingin dilamar:")

	for i = 0; i < jobCount; i++ {
		fmt.Printf("%d. %s di %s\n", i+1, jobs[i].Title, jobs[i].Perusahaan)
	}

	fmt.Print("Pilihan Anda (0 untuk kembali): ")
	var jobIndex int
	fmt.Scan(&jobIndex)

	if jobIndex == 0 {
		return
	}

	if jobIndex < 1 || jobIndex > jobCount {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	jobIndex -= 1

	hitungRelevansiPekerjaan(profile, &jobs[jobIndex])

	var pilihan int
	fmt.Print("\nApa yang ingin Anda buat?",
		" \n1. Resume",
		" \n2. Surat Lamaran",
		" \n3. Keduanya",
		" \nPilihan Anda: ",
	)
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		buatResume(*profile, jobs[jobIndex])
	} else if pilihan == 2 {
		buatSuratLamaran(*profile, jobs[jobIndex])
	} else if pilihan == 3 {
		buatResume(*profile, jobs[jobIndex])
		buatSuratLamaran(*profile, jobs[jobIndex])
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}

func buatResume(profile UserProfile, job Job) {
	var i int

	fmt.Println("\n=====================================")
	fmt.Println("         RESUME")
	fmt.Println("=====================================")

	fmt.Println("\nINFORMASI PRIBADI")
	fmt.Println("-------------------------------------")
	fmt.Printf("Nama    : %s\n", profile.Nama)
	fmt.Printf("Email   : %s\n", profile.Email)
	fmt.Printf("Telepon : %s\n", profile.Phone)
	fmt.Printf("Alamat  : %s\n", profile.Alamat)

	fmt.Println("\nPENGALAMAN KERJA")
	fmt.Println("-------------------------------------")
	if profile.WorkExpCount == 0 {
		fmt.Println("Belum ada pengalaman kerja.")
	} else {
		for i := 0; i < profile.WorkExpCount; i++ {
			exp := profile.PengalamanKerja[i]
			fmt.Printf("%s, %s (%d-%d)\n", exp.Position, exp.Perusahaan, exp.StartYear, exp.EndYear)
			fmt.Printf("  %s\n", exp.Deskripsi)
		}
	}

	fmt.Println("\nPENDIDIKAN")
	fmt.Println("-------------------------------------")
	if profile.EduCount == 0 {
		fmt.Println("Belum ada pendidikan.")
	} else {
		for i = 0; i < profile.EduCount; i++ {
			fmt.Printf("%s %s, %s (%d)\n", profile.Education[i].Degree, profile.Education[i].Bidang, profile.Education[i].Institusi, profile.Education[i].TahunLulus)
		}
	}

	fmt.Println("\nKETERAMPILAN")
	fmt.Println("-------------------------------------")
	if profile.SkillCount == 0 {
		fmt.Println("Belum ada keterampilan.")
	} else {
		for i = 0; i < profile.SkillCount; i++ {
			fmt.Printf("%s (Level %d)\n", profile.Skills[i].Nama, profile.Skills[i].Level)
		}
	}

	fmt.Println("\n*** SARAN AI ***")
	suggestResumeImprovements(profile, job)

	fmt.Println("\nSaya siap saat Anda siap. Tekan Enter untuk melanjutkan...")
	var input string
	fmt.Scanln(&input)
}

func buatSuratLamaran(profile UserProfile, job Job) {
	var now string

	fmt.Println("\n=====================================")
	fmt.Println("         SURAT LAMARAN")
	fmt.Println("======================================")

	fmt.Printf("\n%s\n", profile.Nama)
	fmt.Printf("%s\n", profile.Email)
	fmt.Printf("%s\n", profile.Phone)
	fmt.Printf("%s\n", profile.Alamat)

	now = "18-05-2025"
	fmt.Printf("\n%s\n\n", now)

	fmt.Printf("Kepada Yth.\nHRD %s\n", job.Perusahaan)

	fmt.Printf("\nDengan hormat,\n\n")
	fmt.Printf("Saya bermaksud untuk melamar posisi %s di %s. ", job.Title, job.Perusahaan)
	fmt.Printf("Dengan pengalaman dan keterampilan yang relevan dengan posisi tersebut, saya percaya dapat memberikan kontribusi yang berharga bagi perusahaan Anda.\n\n")

	if profile.WorkExpCount > 0 {
		fmt.Printf("Sebagai seorang yang pernah bekerja sebagai %s di %s, ",
			profile.PengalamanKerja[0].Position, profile.PengalamanKerja[0].Perusahaan)
		fmt.Printf("saya telah mengembangkan kemampuan dalam %s. ",
			profile.PengalamanKerja[0].Deskripsi)
	}

	if profile.SkillCount > 0 {
		fmt.Printf("Saya memiliki keahlian dalam %s ", profile.Skills[0].Nama)
		if profile.SkillCount > 1 {
			fmt.Printf("dan %s ", profile.Skills[1].Nama)
		}
		fmt.Printf("yang akan mendukung performa saya dalam posisi ini.\n\n")
	} else {
		fmt.Printf("\n\n")
	}

	fmt.Printf("Saya sangat tertarik dengan posisi ini karena sesuai dengan minat dan tujuan karir saya. ")
	fmt.Printf("Terima kasih atas perhatian Anda, dan saya berharap dapat bertemu dalam kesempatan wawancara.\n\n")

	fmt.Printf("Hormat saya,\n\n%s\n", profile.Nama)

	fmt.Println("\n*** SARAN AI ***")
	suggestCoverLetterImprovements(job)

	fmt.Println("\nSaya siap saat Anda siap. Tekan Enter untuk melanjutkan...")
	var input string
	fmt.Scanln(&input)
}

func suggestResumeImprovements(profile UserProfile, job Job) {
	var i, j int
	var foundKeywords int
	var keywordFound bool

	fmt.Printf("Tingkat kesesuaian resume Anda dengan pekerjaan: %d%%\n", job.Relevance)

	if job.Relevance < 30 {
		fmt.Println("Saran:")
		fmt.Printf("1. Resume Anda perlu penambahan keterampilan yang terkait dengan kata kunci: ")
		for i = 0; i < job.KeywordCount; i++ {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(job.Keywords[i])
		}
		fmt.Println()
		fmt.Println("2. Tambahkan pengalaman kerja yang lebih relevan dengan posisi ini.")
		fmt.Println("3. Gunakan kata kunci dari deskripsi pekerjaan dalam deskripsi pengalaman Anda.")
	} else if job.Relevance < 70 {
		fmt.Println("Saran:")
		fmt.Println("1. Resume Anda cukup baik, namun bisa ditingkatkan dengan menambahkan lebih banyak detail spesifik.")
		fmt.Printf("2. Pertimbangkan untuk menambahkan keterampilan terkait: ")

		for i = 0; i < job.KeywordCount && foundKeywords < 3; i++ {
			keywordFound = false
			for j = 0; j < profile.SkillCount; j++ {
				if mengandungString(profile.Skills[j].Nama, job.Keywords[i]) {
					keywordFound = true
				}
			}

			if !keywordFound {
				if foundKeywords > 0 {
					fmt.Print(", ")
				}
				fmt.Print(job.Keywords[i])
				foundKeywords++
			}
		}
		fmt.Println()
	} else {
		fmt.Println("Saran:")
		fmt.Println("1. Resume Anda sangat sesuai dengan posisi ini!")
		fmt.Println("2. Pastikan untuk menyoroti pencapaian spesifik untuk menguatkan aplikasi Anda.")
		fmt.Println("3. Pertimbangkan untuk menambahkan contoh proyek atau hasil kerja konkret.")
	}
}

func suggestCoverLetterImprovements(job Job) {
	fmt.Println("Saran untuk surat lamaran:")

	if job.Relevance < 50 {
		fmt.Println("1. Tambahkan lebih banyak contoh spesifik tentang pengalaman Anda yang relevan.")
		fmt.Println("2. Sebutkan alasan konkret mengapa Anda tertarik dengan posisi ini.")
		fmt.Printf("3. Sisipkan kata kunci dari deskripsi pekerjaan seperti: %s", job.Keywords[0])
		if job.KeywordCount > 1 {
			fmt.Printf(", %s", job.Keywords[1])
		}
		fmt.Println()
	} else {
		fmt.Println("1. Surat lamaran Anda sudah baik, namun bisa menyebutkan contoh pencapaian konkret.")
		fmt.Println("2. Tunjukkan pengetahuan Anda tentang perusahaan untuk menunjukkan ketertarikan yang sungguh-sungguh.")
		fmt.Println("3. Sebutkan bagaimana keahlian Anda dapat membantu menyelesaikan masalah spesifik di perusahaan.")
	}
}

func cariPekerjaan(jobs tabJob, jobCount int) {
	var isSearching bool = true
	for isSearching {
		fmt.Print("\n*** CARI PEKERJAAN ***",
			" \n1. Cari berdasarkan Posisi/Jabatan (Sequential Search)",
			" \n2. Cari berdasarkan Kata Kunci Industri (Binary Search)",
			" \n0. Kembali",
			" \nPilihan Anda: ",
		)

		var pilihan int = -1
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			sequentialSearch(jobs, jobCount)
		} else if pilihan == 2 {
			pencarianBiner(jobs, jobCount)
		} else if pilihan == 0 {
			isSearching = false
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}

	}
}

func sequentialSearch(jobs tabJob, jobCount int) {
	var i int
	var found bool = false

	fmt.Print("\nMasukkan kata kunci posisi yang dicari: ")
	var keyword string
	fmt.Scan(&keyword)

	fmt.Println("\nHasil pencarian:")
	fmt.Println("-------------------------------------")

	for i = 0; i < jobCount; i++ {
		if mengandungString(jobs[i].Title, keyword) {
			fmt.Printf("%d. %s di %s\n", i+1, jobs[i].Title, jobs[i].Perusahaan)
			fmt.Printf("   Deskripsi: %s\n", jobs[i].Deskripsi)
			fmt.Printf("   Gaji: Rp %d\n", jobs[i].Salary)
			fmt.Println("-------------------------------------")
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ditemukan pekerjaan yang sesuai dengan kata kunci tersebut.")
	}

	fmt.Println("\nSaya siap saat Anda siap. Tekan Enter untuk melanjutkan...")
	var input string
	fmt.Scanln(&input)
}

func pencarianBiner(jobs tabJob, jobCount int) {
	var i, j, low, high, mid int
	var found bool = false
	var searchKeyword string
	for i = 0; i < jobCount-1; i++ {
		for j = 0; j < jobCount-i-1; j++ {
			if jobs[j].KeywordCount > 0 && jobs[j+1].KeywordCount > 0 {
				if jobs[j].Keywords[0] > jobs[j+1].Keywords[0] {
					jobs[j], jobs[j+1] = jobs[j+1], jobs[j]
				}
			}
		}
	}

	fmt.Println("\nKata Kunci Industri yang Tersedia:")
	uniqueKeywords := [20]string{}
	uniqueCount := 0

	for i = 0; i < jobCount; i++ {
		if jobs[i].KeywordCount > 0 {
			keywordExists := false
			for j = 0; j < uniqueCount; j++ {
				if jobs[i].Keywords[0] == uniqueKeywords[j] {
					keywordExists = true
				}
			}

			if !keywordExists && uniqueCount < 20 {
				uniqueKeywords[uniqueCount] = jobs[i].Keywords[0]
				uniqueCount += 1
			}
		}
	}

	for i = 0; i < uniqueCount; i++ {
		fmt.Printf("%d. %s\n", i+1, uniqueKeywords[i])
	}

	fmt.Print("\nPilih nomor kata kunci industri: ")
	var keywordIndex int
	fmt.Scan(&keywordIndex)

	if keywordIndex < 1 || keywordIndex > uniqueCount {
		fmt.Println("Pilihan tidak valid.")
		return
	}

	searchKeyword = uniqueKeywords[keywordIndex-1]

	low = 0
	high = jobCount - 1
	found = false

	for low <= high {
		mid = (low + high) / 2

		if jobs[mid].KeywordCount > 0 && jobs[mid].Keywords[0] == searchKeyword {
			found = true

			fmt.Println("\nHasil pencarian:")
			fmt.Println("-------------------------------------")

			for i := mid; i >= 0; i-- {
				if jobs[i].KeywordCount > 0 && jobs[i].Keywords[0] == searchKeyword {
					fmt.Printf("%s di %s\n", jobs[i].Title, jobs[i].Perusahaan)
					fmt.Printf("Deskripsi: %s\n", jobs[i].Deskripsi)
					fmt.Printf("Gaji: Rp %d\n", jobs[i].Salary)
					fmt.Println("-------------------------------------")
				} else if i < mid {
					i = -1
				}
			}

			for i = mid + 1; i < jobCount; i++ {
				if jobs[i].KeywordCount > 0 && jobs[i].Keywords[0] == searchKeyword {
					fmt.Printf("%s di %s\n", jobs[i].Title, jobs[i].Perusahaan)
					fmt.Printf("Deskripsi: %s\n", jobs[i].Deskripsi)
					fmt.Printf("Gaji: Rp %d\n", jobs[i].Salary)
					fmt.Println("-------------------------------------")
				} else {
					i = jobCount
				}
			}

			low = high + 1
		} else if jobs[mid].KeywordCount > 0 && jobs[mid].Keywords[0] < searchKeyword {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !found {
		fmt.Println("Tidak ditemukan pekerjaan yang sesuai dengan kata kunci tersebut.")
	}

	fmt.Println("\nSaya siap saat Anda siap. Tekan Enter untuk melanjutkan...")
	var input string
	fmt.Scanln(&input)
}

func urutkanPekerjaan(jobs tabJob, jobCount int) {
	var isSorting bool = true
	for isSorting {

		fmt.Print("\n*** URUTKAN DAFTAR PEKERJAAN ***",
			" \n1. Urutkan berdasarkan Relevansi (Selection Sort)",
			" \n2. Urutkan berdasarkan Gaji (Insertion Sort)",
			" \n0. Kembali",
			" \nPilihan Anda: ",
		)

		var pilihan int = -1
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			selectionSort(jobs, jobCount)
		} else if pilihan == 2 {
			insertionSort(jobs, jobCount)
		} else if pilihan == 0 {
			isSorting = false
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func selectionSort(jobs tabJob, jobCount int) {
	var i, idx, pass int
	var temp Job

	pass = 1
	for pass < jobCount {
		idx = pass - 1
		i = pass
		for i < jobCount {
			if jobs[i].Relevance > jobs[idx].Relevance {
				idx = i
			}
			i = i + 1
		}

		temp = jobs[pass-1]
		jobs[pass-1] = jobs[idx]
		jobs[idx] = temp

		pass = pass + 1
	}

	fmt.Println("\nDaftar Pekerjaan berdasarkan Relevansi (Tertinggi ke Terendah):")
	fmt.Println("----------------------------------------------------------------")

	for i := 0; i < jobCount; i++ {
		fmt.Printf("%d. %s di %s\n", i+1, jobs[i].Title, jobs[i].Perusahaan)
		fmt.Printf("   Relevansi: %d%%\n", jobs[i].Relevance)
		fmt.Printf("   Gaji: Rp %d\n", jobs[i].Salary)
		fmt.Println("----------------------------------------------------------------")
	}

	fmt.Println("\nSaya siap saat Anda siap. Tekan Enter untuk melanjutkan...")
	var input string
	fmt.Scanln(&input)
}

func insertionSort(jobs tabJob, jobCount int) {
	var i, pass int
	var temp Job

	pass = 1
	for pass <= jobCount-1 {
		i = pass
		temp = jobs[pass]
		for i > 0 && temp.Salary > jobs[i-1].Salary {
			jobs[i] = jobs[i-1]
			i = i - 1
		}
		jobs[i] = temp
		pass = pass + 1
	}

	fmt.Println("\nDaftar Pekerjaan berdasarkan Gaji (Tertinggi ke Terendah):")
	fmt.Println("----------------------------------------------------------------")

	for i = 0; i < jobCount; i++ {
		fmt.Printf("%d. %s di %s\n", i+1, jobs[i].Title, jobs[i].Perusahaan)
		fmt.Printf("   Gaji: Rp %d\n", jobs[i].Salary)
		fmt.Printf("   Relevansi: %d%%\n", jobs[i].Relevance)
		fmt.Println("----------------------------------------------------------------")
	}

	fmt.Println("\nSaya siap saat Anda siap. Tekan Enter untuk melanjutkan...")
	var input string
	fmt.Scanln(&input)
}

func evaluasiResume(profile *UserProfile, jobs [20]Job, jobCount int) {
	var i, j int
	var keyword string
	var keywordFound bool
	var foundMissingKeywords bool

	if profile.Nama == "" || profile.WorkExpCount == 0 && profile.SkillCount == 0 {
		fmt.Println("\nAnda perlu melengkapi profil terlebih dahulu.")
		return
	}

	for i = 0; i < jobCount; i++ {
		hitungRelevansiPekerjaan(profile, &jobs[i])
	}

	var mostRelevantIndex int = 0
	for i = 1; i < jobCount; i++ {
		if jobs[i].Relevance > jobs[mostRelevantIndex].Relevance {
			mostRelevantIndex = i
		}
	}

	fmt.Println("\n*** EVALUASI RESUME ***")
	fmt.Println("----------------------------------------------------------------")

	fmt.Printf("Nama: %s\n", profile.Nama)
	fmt.Printf("Jumlah Pengalaman Kerja: %d\n", profile.WorkExpCount)
	fmt.Printf("Jumlah Keterampilan: %d\n", profile.SkillCount)
	fmt.Printf("Jumlah Pendidikan: %d\n", profile.EduCount)

	fmt.Println("\nPekerjaan Paling Cocok:")
	fmt.Printf("%s di %s (Relevansi: %d%%)\n",
		jobs[mostRelevantIndex].Title,
		jobs[mostRelevantIndex].Perusahaan,
		jobs[mostRelevantIndex].Relevance)

	fmt.Println("\nKekuatan Resume:")

	if profile.SkillCount > 0 {
		var topSkillIndex int = 0
		for i = 1; i < profile.SkillCount; i++ {
			if profile.Skills[i].Level > profile.Skills[topSkillIndex].Level {
				topSkillIndex = i
			}
		}

		fmt.Printf("1. Keahlian utama: %s (Level %d)\n",
			profile.Skills[topSkillIndex].Nama,
			profile.Skills[topSkillIndex].Level)
	}

	if profile.WorkExpCount > 0 {
		var latestExpIndex int = 0
		for i = 1; i < profile.WorkExpCount; i++ {
			if profile.PengalamanKerja[i].StartYear > profile.PengalamanKerja[latestExpIndex].StartYear {
				latestExpIndex = i
			}
		}

		fmt.Printf("2. Pengalaman terbaru: %s di %s\n",
			profile.PengalamanKerja[latestExpIndex].Position,
			profile.PengalamanKerja[latestExpIndex].Perusahaan)
	}

	if profile.EduCount > 0 {
		var latestEduIndex int = 0
		for i = 1; i < profile.EduCount; i++ {
			if profile.Education[i].TahunLulus > profile.Education[latestEduIndex].TahunLulus {
				latestEduIndex = i
			}
		}

		fmt.Printf("3. Pendidikan terbaru: %s %s dari %s\n",
			profile.Education[latestEduIndex].Degree,
			profile.Education[latestEduIndex].Bidang,
			profile.Education[latestEduIndex].Institusi)
	}

	fmt.Println("\nSaran Peningkatan:")

	if profile.WorkExpCount < 2 {
		fmt.Println("1. Tambahkan lebih banyak pengalaman kerja untuk memperkuat resume.")
	}

	if profile.SkillCount < 5 {
		fmt.Println("2. Tambahkan lebih banyak keterampilan untuk meningkatkan daya saing.")
	}

	if jobCount > 0 {
		fmt.Println("3. Pertimbangkan untuk menambahkan keterampilan berikut:")

		foundMissingKeywords = false
		for i = 0; i < jobs[mostRelevantIndex].KeywordCount; i++ {
			keyword = jobs[mostRelevantIndex].Keywords[i]

			keywordFound = false
			for j = 0; j < profile.SkillCount; j++ {
				if mengandungString(profile.Skills[j].Nama, keyword) {
					keywordFound = true
				}
			}

			if !keywordFound {
				fmt.Printf("   - %s\n", keyword)
				foundMissingKeywords = true
			}
		}

		if !foundMissingKeywords {
			fmt.Println("   - Keterampilan Anda sudah sesuai dengan pekerjaan target!")
		}
	}

	fmt.Println("\nSaya siap saat Anda siap. Tekan Enter untuk melanjutkan...")
	var input string
	fmt.Scanln(&input)
}
