-- +goose Up
INSERT INTO users (name, email, hashed_password) VALUES (
		'Test User',
		'test@example.com',
		'$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG'
);

-- +goose Down
