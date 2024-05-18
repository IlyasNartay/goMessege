package proto

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/lib/pq" // Импорт PostgreSQL драйвера для database/sql
)

type Server struct {
	Db *sql.DB
}

func (s Server) RegisterUser(ctx context.Context, in *RegisterUserRequest) (*RegisterUserResponse, error) {
	username := in.GetUsername()
	password := in.GetPassword()

	// Проверяем, существует ли уже пользователь с таким именем пользователя или email в базе данных
	var count int
	err := s.Db.QueryRowContext(ctx, "SELECT COUNT(*) FROM users WHERE username=$1", username).Scan(&count)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("user with this username or email already exists")
	}
	//
	//// Хэшируем пароль перед сохранением
	//// В реальном приложении используйте хэширование и соль для пароля, а не храните его в чистом виде
	//// Здесь просто для демонстрации
	//hashedPassword := hashPassword(password)

	// Вставляем нового пользователя в базу данных
	_, err = s.Db.ExecContext(ctx, "INSERT INTO users (username, email, password) VALUES ($1, $2)", username, password)
	if err != nil {
		return nil, err
	}

	// Возвращаем успешный ответ
	return &RegisterUserResponse{
		Message: "User registered successfully",
	}, nil
}

func (s Server) Login(ctx context.Context, in *LoginRequest) (*LoginResponse, error) {
	username := in.GetUsername()
	//password := in.GetPassword()

	// Получаем пользователя из базы данных по имени пользователя
	var storedPassword string
	err := s.Db.QueryRowContext(ctx, "SELECT password FROM users WHERE username=$1", username).Scan(&storedPassword)
	if err != nil {
		return nil, err
	}

	// Проверяем правильность пароля
	//if !checkPassword(password, storedPassword) {
	//	return nil, errors.New("invalid username or password")
	//}

	// Возвращаем успешный ответ
	return &LoginResponse{
		Token: "new Token",
	}, nil
}

func (s Server) RefreshToken(ctx context.Context, request *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	return &RefreshTokenResponse{
		Token: "new_access_token",
	}, nil
}

func (s Server) mustEmbedUnimplementedAuthServiceServer() {

}

//
//func (s server) RefreshToken(ctx context.Context, in *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
//	// Ваша логика обновления токена доступа
//	// В данном примере просто возвращаем новый токен

// }
func hashPassword(password string) string {
	// В реальном приложении реализуйте хэширование пароля, например, с использованием bcrypt
	// Здесь просто для демонстрации
	return password
}

func checkPassword(password, hashedPassword string) bool {
	// В реальном приложении реализуйте проверку пароля
	// Здесь просто для демонстрации
	return password == hashedPassword
}
