package seeders

import (
	"log"
	"time"

	"github.com/blog-cms/backend/models"
	"gorm.io/gorm"
)

func SeedData(db *gorm.DB) {
	// Cek apakah data sudah ada
	var authorCount int64
	db.Model(&models.Author{}).Count(&authorCount)
	if authorCount > 0 {
		log.Println("📦 Data sudah ada, melewati seeder")
		return
	}

	log.Println("🌱 Memulai seeding data...")

	// === SEED AUTHORS ===
	authors := []models.Author{
		{
			Name:   "Ahmad Fauzi",
			Email:  "ahmad.fauzi@blog.com",
			Bio:    "Full-stack developer dan penulis teknis dengan pengalaman lebih dari 5 tahun dalam pengembangan web. Suka berbagi pengetahuan tentang Go, React, dan arsitektur software.",
			Avatar: "https://ui-avatars.com/api/?name=Ahmad+Fauzi&background=6366f1&color=fff&size=200",
		},
		{
			Name:   "Siti Nurhaliza",
			Email:  "siti.nurhaliza@blog.com",
			Bio:    "UI/UX Designer dan frontend developer yang passionate tentang desain interaktif dan pengalaman pengguna. Menulis tentang CSS, design systems, dan best practices.",
			Avatar: "https://ui-avatars.com/api/?name=Siti+Nurhaliza&background=ec4899&color=fff&size=200",
		},
		{
			Name:   "Budi Santoso",
			Email:  "budi.santoso@blog.com",
			Bio:    "DevOps engineer dan cloud architect. Menulis tentang containerization, CI/CD, dan infrastruktur cloud modern.",
			Avatar: "https://ui-avatars.com/api/?name=Budi+Santoso&background=10b981&color=fff&size=200",
		},
		{
			Name:   "Dewi Lestari",
			Email:  "dewi.lestari@blog.com",
			Bio:    "Data scientist dan AI enthusiast. Menulis tentang machine learning, data analysis, dan Python.",
			Avatar: "https://ui-avatars.com/api/?name=Dewi+Lestari&background=f59e0b&color=fff&size=200",
		},
	}

	for i := range authors {
		db.Create(&authors[i])
	}
	log.Println("✅ Seeded 4 authors")

	// === SEED CATEGORIES ===
	categories := []models.Category{
		{Name: "Teknologi", Description: "Artikel tentang perkembangan teknologi terkini, gadget, dan inovasi digital", Color: "#6366F1"},
		{Name: "Pemrograman", Description: "Tutorial dan tips pemrograman berbagai bahasa dan framework", Color: "#EC4899"},
		{Name: "Desain", Description: "Artikel tentang UI/UX design, graphic design, dan creative tools", Color: "#10B981"},
		{Name: "Cloud & DevOps", Description: "Artikel tentang cloud computing, containerization, dan CI/CD", Color: "#F59E0B"},
		{Name: "Data Science", Description: "Artikel tentang machine learning, AI, dan analisis data", Color: "#EF4444"},
		{Name: "Karier", Description: "Tips dan panduan untuk pengembangan karier di industri teknologi", Color: "#8B5CF6"},
	}

	for i := range categories {
		db.Create(&categories[i])
	}
	log.Println("✅ Seeded 6 categories")

	// === SEED ARTICLES ===
	now := time.Now()
	articles := []models.Article{
		{
			Title:       "Memulai Pengembangan API dengan Golang dan Gin Framework",
			Slug:        "memulai-pengembangan-api-dengan-golang-dan-gin-framework",
			Content:     "Go (atau Golang) adalah bahasa pemrograman yang dikembangkan oleh Google. Bahasa ini terkenal dengan performa tinggi, concurrency yang powerful, dan syntax yang sederhana. Gin adalah web framework untuk Go yang sangat cepat dan minimalis.\n\nDalam artikel ini, kita akan membahas cara membangun RESTful API menggunakan Gin Framework. Mulai dari setup project, membuat routes, controllers, hingga integrasi dengan database PostgreSQL menggunakan GORM.\n\n## Mengapa Gin Framework?\n\nGin menyediakan performa yang sangat baik dibandingkan framework lain. Beberapa kelebihannya:\n- Routing yang cepat dengan httprouter\n- Middleware support\n- JSON validation\n- Error management\n- Rendering built-in\n\n## Setup Project\n\nLangkah pertama adalah membuat module Go dan menginstall dependencies yang diperlukan. Gunakan `go mod init` untuk memulai project baru.\n\n## Kesimpulan\n\nGin Framework adalah pilihan yang tepat untuk membangun API yang cepat dan efisien dengan Go. Dengan arsitektur MVC dan GORM sebagai ORM, kita bisa membangun aplikasi yang scalable dan maintainable.",
			Excerpt:     "Panduan lengkap membangun RESTful API dengan Go dan Gin Framework. Mulai dari setup hingga deployment.",
			CoverImage:  "https://images.unsplash.com/photo-1555066931-4365d14bab8c?w=800",
			Status:      "published",
			AuthorID:    1,
			CategoryID:  2,
			PublishedAt: &now,
		},
		{
			Title:       "React + Vite: Kombinasi Modern untuk Frontend Development",
			Slug:        "react-vite-kombinasi-modern-frontend-development",
			Content:     "Vite adalah build tool generasi terbaru yang menawarkan pengalaman development yang sangat cepat. Dikombinasikan dengan React, kita mendapatkan workflow development yang sangat produktif.\n\n## Mengapa Vite?\n\nVite menggunakan native ES modules dan menawarkan Hot Module Replacement (HMR) yang sangat cepat. Dibandingkan webpack, Vite:\n- Startup time yang sangat cepat\n- HMR yang instan\n- Build yang optimal dengan Rollup\n- Support TypeScript out of the box\n\n## Setup React dengan Vite\n\nMembuat project React dengan Vite sangat mudah. Cukup jalankan `npm create vite@latest` dan pilih template React.\n\n## Component Architecture\n\nDalam React modern, kita menggunakan functional components dan hooks. Pattern yang umum digunakan:\n- Custom hooks untuk logic reusable\n- Context API untuk state management\n- React Router untuk routing\n\n## Kesimpulan\n\nReact + Vite adalah kombinasi yang powerful untuk membangun aplikasi frontend modern. Dengan ekosistem yang kaya dan performa yang excellent, ini adalah pilihan yang tepat untuk project apapun.",
			Excerpt:     "Pelajari cara menggunakan React dengan Vite untuk pengalaman development yang super cepat dan modern.",
			CoverImage:  "https://images.unsplash.com/photo-1633356122544-f134324a6cee?w=800",
			Status:      "published",
			AuthorID:    1,
			CategoryID:  2,
			PublishedAt: &now,
		},
		{
			Title:       "Prinsip UI/UX Design yang Harus Diketahui Setiap Developer",
			Slug:        "prinsip-ui-ux-design-untuk-developer",
			Content:     "Sebagai developer, memahami prinsip dasar UI/UX design bisa sangat membantu dalam membangun produk yang lebih baik. Tidak hanya membuat tampilan cantik, tapi juga memastikan pengalaman pengguna yang optimal.\n\n## Hierarchy Visual\n\nHierarchy visual membantu pengguna memahami konten dengan cepat. Gunakan ukuran, warna, dan spacing untuk membedakan elemen penting dari yang kurang penting.\n\n## Konsistensi\n\nDesain yang konsisten membuat aplikasi terasa profesional dan mudah dipelajari. Gunakan design system atau component library untuk menjaga konsistensi.\n\n## Accessibility\n\nPastikan aplikasi bisa digunakan oleh semua orang, termasuk pengguna dengan disabilitas. Gunakan semantic HTML, ARIA labels, dan pastikan contrast ratio yang memadai.\n\n## Responsive Design\n\nDesain harus bekerja dengan baik di semua ukuran layar. Gunakan flexbox, grid, dan media queries untuk membuat layout yang responsif.",
			Excerpt:     "Pahami prinsip dasar UI/UX design untuk membangun produk digital yang lebih baik dan user-friendly.",
			CoverImage:  "https://images.unsplash.com/photo-1561070791-2526d30994b5?w=800",
			Status:      "published",
			AuthorID:    2,
			CategoryID:  3,
			PublishedAt: &now,
		},
		{
			Title:       "Docker dan Kubernetes: Panduan Kontainerisasi untuk Pemula",
			Slug:        "docker-kubernetes-panduan-kontainerisasi-pemula",
			Content:     "Kontainerisasi telah mengubah cara kita membangun, mendistribusikan, dan menjalankan aplikasi. Docker dan Kubernetes adalah dua teknologi kunci dalam ekosistem container.\n\n## Apa itu Docker?\n\nDocker memungkinkan kita untuk mengemas aplikasi beserta dependensinya ke dalam container yang portabel. Container ini bisa dijalankan di mana saja dengan konsisten.\n\n## Apa itu Kubernetes?\n\nKubernetes (K8s) adalah platform orchestration untuk container. Kubernetes mengelola deployment, scaling, dan operasi container secara otomatis.\n\n## Docker Compose\n\nUntuk menjalankan multiple container secara bersamaan, kita menggunakan Docker Compose. Ini sangat berguna untuk development environment.\n\n## Best Practices\n\n- Gunakan multi-stage builds\n- Minimize image size\n- Jangan jalankan container sebagai root\n- Gunakan health checks",
			Excerpt:     "Pelajari dasar-dasar Docker dan Kubernetes untuk memulai perjalanan kontainerisasi Anda.",
			CoverImage:  "https://images.unsplash.com/photo-1667372393119-3d4c48d07fc9?w=800",
			Status:      "published",
			AuthorID:    3,
			CategoryID:  4,
			PublishedAt: &now,
		},
		{
			Title:       "Pengenalan Machine Learning dengan Python",
			Slug:        "pengenalan-machine-learning-python",
			Content:     "Machine Learning (ML) adalah cabang dari Artificial Intelligence yang memungkinkan komputer untuk belajar dari data tanpa diprogram secara eksplisit. Python adalah bahasa yang paling populer untuk ML.\n\n## Library Penting\n\n- NumPy: Komputasi numerik\n- Pandas: Manipulasi data\n- Scikit-learn: Algoritma ML\n- TensorFlow/PyTorch: Deep Learning\n\n## Jenis Machine Learning\n\n1. Supervised Learning\n2. Unsupervised Learning\n3. Reinforcement Learning\n\n## Workflow ML\n\n1. Data Collection\n2. Data Preprocessing\n3. Feature Engineering\n4. Model Training\n5. Model Evaluation\n6. Deployment",
			Excerpt:     "Mulai perjalanan Machine Learning Anda dengan Python. Panduan lengkap dari dasar hingga implementasi.",
			CoverImage:  "https://images.unsplash.com/photo-1515879218367-8466d910auj7?w=800",
			Status:      "published",
			AuthorID:    4,
			CategoryID:  5,
			PublishedAt: &now,
		},
		{
			Title:       "Tips Membangun Karier di Industri Teknologi",
			Slug:        "tips-membangun-karier-industri-teknologi",
			Content:     "Industri teknologi terus berkembang dan menawarkan banyak peluang karier yang menarik. Berikut adalah tips untuk membangun karier yang sukses.\n\n## Bangun Portfolio\n\nPortfolio yang kuat lebih penting dari sertifikat. Bangun proyek nyata yang menunjukkan kemampuan Anda.\n\n## Networking\n\nBergabung dengan komunitas developer, ikuti meetup, dan aktif di platform seperti GitHub dan Stack Overflow.\n\n## Terus Belajar\n\nTeknologi berubah dengan cepat. Dedikasikan waktu untuk mempelajari hal baru setiap hari.\n\n## Soft Skills\n\nJangan abaikan soft skills seperti komunikasi, teamwork, dan problem solving. Ini sama pentingnya dengan technical skills.",
			Excerpt:     "Panduan praktis untuk membangun karier yang sukses di industri teknologi yang kompetitif.",
			CoverImage:  "https://images.unsplash.com/photo-1522202176988-66273c2fd55f?w=800",
			Status:      "draft",
			AuthorID:    2,
			CategoryID:  6,
		},
		{
			Title:       "PostgreSQL vs MySQL: Memilih Database yang Tepat",
			Slug:        "postgresql-vs-mysql-memilih-database-tepat",
			Content:     "Memilih database yang tepat adalah keputusan penting dalam pengembangan software. PostgreSQL dan MySQL adalah dua database relasional yang paling populer.\n\n## PostgreSQL\n\nPostgreSQL dikenal sebagai database yang paling advanced. Kelebihannya:\n- Support JSON/JSONB\n- Full-text search bawaan\n- Extensible\n- ACID compliant\n\n## MySQL\n\nMySQL adalah database yang paling banyak digunakan di dunia. Kelebihannya:\n- Mudah digunakan\n- Performa read yang cepat\n- Komunitas besar\n- Banyak hosting support\n\n## Kapan Menggunakan Apa?\n\nGunakan PostgreSQL untuk aplikasi yang membutuhkan fitur advanced dan data integrity yang ketat. Gunakan MySQL untuk aplikasi yang membutuhkan simplisitas dan performa read tinggi.",
			Excerpt:     "Perbandingan mendalam antara PostgreSQL dan MySQL untuk membantu Anda memilih database yang tepat.",
			CoverImage:  "https://images.unsplash.com/photo-1544383835-bda2bc66a55d?w=800",
			Status:      "published",
			AuthorID:    3,
			CategoryID:  1,
			PublishedAt: &now,
		},
	}

	for i := range articles {
		db.Create(&articles[i])
	}
	log.Println("✅ Seeded 7 articles")

	// === SEED COMMENTS ===
	comments := []models.Comment{
		{ArticleID: 1, CommenterName: "Rizky Pratama", CommenterEmail: "rizky@email.com", Content: "Artikel yang sangat membantu! Saya berhasil membangun API pertama saya dengan Gin. Terima kasih!"},
		{ArticleID: 1, CommenterName: "Dina Safitri", CommenterEmail: "dina@email.com", Content: "Bisa tolong tambahkan contoh dengan authentication middleware? Itu akan sangat berguna."},
		{ArticleID: 1, CommenterName: "Andi Wijaya", CommenterEmail: "andi@email.com", Content: "Clean dan jelas penjelasannya. Bookmarked untuk referensi!"},
		{ArticleID: 2, CommenterName: "Maya Putri", CommenterEmail: "maya@email.com", Content: "Saya sudah pindah dari CRA ke Vite dan perbedaan kecepatannya luar biasa! Highly recommended."},
		{ArticleID: 2, CommenterName: "Fajar Nugroho", CommenterEmail: "fajar@email.com", Content: "Apakah Vite support SSR juga? Atau harus pakai framework seperti Next.js?"},
		{ArticleID: 3, CommenterName: "Lina Kusuma", CommenterEmail: "lina@email.com", Content: "Sebagai developer yang sering lupa tentang accessibility, artikel ini sangat penting. Terima kasih sudah mengingatkan!"},
		{ArticleID: 3, CommenterName: "Hendra Saputra", CommenterEmail: "hendra@email.com", Content: "Tips tentang visual hierarchy sangat berguna. Sudah saya terapkan di project saya."},
		{ArticleID: 4, CommenterName: "Rahma Dewi", CommenterEmail: "rahma@email.com", Content: "Penjelasan Docker-nya sangat mudah dipahami. Bisakah dibuatkan tutorial untuk Docker Compose juga?"},
		{ArticleID: 5, CommenterName: "Yusuf Hakim", CommenterEmail: "yusuf@email.com", Content: "Artikel yang bagus untuk pemula seperti saya. Sekarang saya lebih paham workflow ML."},
		{ArticleID: 7, CommenterName: "Sari Indah", CommenterEmail: "sari@email.com", Content: "Perbandingan yang sangat objektif. Akhirnya saya memutuskan untuk menggunakan PostgreSQL untuk project baru saya."},
	}

	for i := range comments {
		db.Create(&comments[i])
	}
	log.Println("✅ Seeded 10 comments")

	log.Println("🎉 Seeding data selesai!")
}
