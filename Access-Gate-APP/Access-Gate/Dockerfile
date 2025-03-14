# Etapa 1: Construcción de la aplicación Angular
FROM node:20 as build

WORKDIR /app

# Copia los archivos necesarios y instala dependencias
COPY package.json package-lock.json ./
RUN npm ci

# Copia el código fuente
COPY . .

# Compila la aplicación en modo producción
RUN npm run build -- --configuration=production

# Etapa 2: Servir la aplicación con Nginx
FROM nginx:stable-alpine

# Copia la aplicación compilada desde la etapa anterior
COPY --from=build /app/dist/access-gate /usr/share/nginx/html

# Copia configuración de Nginx
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Exponer el puerto 80
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
