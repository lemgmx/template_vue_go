# template_vue_go

This project combines a Go backend API with a Vue 3 frontend.

## Backend (Go API)

The backend provides a ready-to-use Go API project with a clean architecture structure. It includes built-in support for:

- Database configuration (MySQL)
- Hot reloading with Air
- MVC-like organization (controllers, models, services)
- Separate route definitions
- Example resources (people and planets)

### Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) with the [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.Go).

### Project Setup

1. Install Go (version 1.21 or later)
2. Install Air for hot reloading:
```sh
go install github.com/cosmtrek/air@latest
```
3. Copy `.env.example` to `.env` and configure your database settings:
```sh
cp .env.example .env
```

### Development Commands

#### Run with hot-reload
```sh
air
```

#### Run without hot-reload
```sh
go run server.go
```

### Project Structure

- `controllers/`: Request handlers
- `models/`: Data models
- `services/`: Business logic
- `routes/`: Route definitions
- `database/`: Database configuration and connection

## Frontend (Vue 3)

The frontend is a Vue 3 application built with Vite, featuring:

- Composition API
- Pinia for state management
- Vue Router for navigation
- Tailwind CSS for styling
- Dark/light theme support
- Axios for API communication

### Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) + [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) (and disable Vetur).

### Project Setup

```sh
npm install
```

### Development Commands

#### Compile and Hot-Reload for Development
```sh
npm run dev
```

#### Compile and Minify for Production
```sh
npm run build
```

#### Lint with ESLint
```sh
npm run lint
```

### Project Structure

- `src/`: Main application source
  - `assets/`: Static assets (images, styles)
  - `components/`: Reusable Vue components
  - `router/`: Vue Router configuration
  - `services/`: API service layer
  - `stores/`: Pinia state stores
  - `views/`: Page components
  - `App.vue`: Root component
  - `main.js`: Application entry point

### Key Features

- Pre-configured API service with Axios
- Example state management with Pinia
- Responsive design with Tailwind CSS
- Dark/light theme support
- Clean component architecture
