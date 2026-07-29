package recommendation

import "time"

// Stream representa los datos necesarios para rankear
type Stream struct {
	ID        int
	Views     int
	Likes     int
	CreatedAt time.Time
}

// CalculateScore devuelve un número que representa la relevancia
func CalculateScore(s Stream) float64 {

	score := 0.0

	// 📊 Vistas
	score += float64(s.Views) * 1.5

	// ❤️ Likes (más peso)
	score += float64(s.Likes) * 3

	// 🆕 Bonus por recencia
	hours := time.Since(s.CreatedAt).Hours()

	if hours < 24 {
		score += 50
	} else if hours < 72 {
		score += 20
	}

	return score
}
