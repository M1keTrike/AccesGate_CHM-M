import { Component, OnInit, Inject } from '@angular/core';
import { Messaging, getToken, onMessage } from '@angular/fire/messaging';
import { HttpClient } from '@angular/common/http';
import { UsersService } from '../../../../services/Users.Service';
import { Router } from '@angular/router';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-admin-dashboard',
  templateUrl: './admin-dashboard.component.html',
  styleUrls: ['./admin-dashboard.component.css'],
  standalone: false,
})
export class AdminDashboardComponent implements OnInit {
  token: string = '';
  private backendUrl = 'http://localhost:8085'; // ✅ URL base del backend

  constructor(
    @Inject(Messaging) private messaging: Messaging,
    private http: HttpClient,
    private usersService: UsersService,
    private router: Router
  ) {}

  ngOnInit() {
    this.listenForMessages();
  }

  // ✅ Solicitar permisos y obtener el token de FCM
  requestPermission() {
    Notification.requestPermission()
      .then(async (permission) => {
        if (permission === 'granted') {
          try {
            const token = await getToken(this.messaging, {
              vapidKey:
                'BNiXbBcCoErAiquuylp5PsU2nT8I1Tj4fbX-JPzEj1nyb7A3lQuNxKdZuSy-J4W9QkhPFjT05SQC5s1cv64GlB8',
            });

            if (token) {
              console.log('🔑 Token de notificación:', token);
              this.token = token;

              // ✅ Enviar token al backend para suscribirse al tema "all"
              this.subscribeToBackend(token);
            }
          } catch (err) {
            console.error('🚨 Error obteniendo token de FCM:', err);
          }
        } else {
          console.warn('⚠️ Permiso de notificaciones no concedido');
        }
      })
      .catch((error) => console.error('🚨 Error solicitando permisos:', error));
  }

  // ✅ Enviar token al backend para suscribirse al tema "all"
  private subscribeToBackend(token: string) {
    this.http.post<{ message: string }>(`${this.backendUrl}/subscribe`, { token })
      .subscribe({
        next: (res) => console.log(res.message),
        error: (err) => console.error('🚨 Error en la suscripción:', err),
      });
  }

  // 📩 Escuchar mensajes en primer plano
  listenForMessages() {
    onMessage(this.messaging, (payload) => {
      console.log('📩 Mensaje recibido en primer plano:', payload);
      Swal.fire({
        title: payload.notification?.title || 'Nueva Notificación',
        text: payload.notification?.body || 'Tienes un nuevo mensaje.',
        icon: 'info',
        toast: true,
        position: 'top-end',
        showConfirmButton: false,
        timer: 5000,
        timerProgressBar: true,
      });
    });
  }

  logout() {
    this.usersService.logout();
    this.router.navigate(['/login']);
  }
}
