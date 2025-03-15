import { Component, OnInit, Inject } from '@angular/core';
import { Messaging, getToken, onMessage } from '@angular/fire/messaging';
import Swal from 'sweetalert2';

@Component({
    selector: 'app-admin-dashboard',
    templateUrl: './admin-dashboard.component.html',
    styleUrls: ['./admin-dashboard.component.css'], // Corregido "styleUrls"
    standalone: false
})
export class AdminDashboardComponent implements OnInit {
    token: string = '';

    constructor(@Inject(Messaging) private messaging: Messaging) {}

    ngOnInit() {
        this.listenForMessages();
    }

    requestPermission() {
        Notification.requestPermission().then((permission) => {
            if (permission === 'granted') {
                getToken(this.messaging, {
                    vapidKey: 'BNiXbBcCoErAiquuylp5PsU2nT8I1Tj4fbX-JPzEj1nyb7A3lQuNxKdZuSy-J4W9QkhPFjT05SQC5s1cv64GlB8'
                })
                .then((token) => {
                    if (token) {
                        console.log('Token de notificación:', token);
                        this.token = token;
                    }
                })
                .catch((err) => console.error('Error obteniendo token', err));
            } else {
                console.warn('Permiso de notificaciones no concedido');
            }
        }).catch((error) => console.error('Error solicitando permisos', error));
    }

    listenForMessages() {
        onMessage(this.messaging, (payload) => {
            console.log('Mensaje recibido en primer plano:', payload);
            Swal.fire({
                title: payload.notification?.title || 'Nueva Notificación',
                text: payload.notification?.body || 'Tienes un nuevo mensaje.',
                icon: 'info',
                toast: true,              
                position: 'top-end',      
                showConfirmButton: false, 
                timer: 5000,              
                timerProgressBar: true    
                
              });
        });
    }
}
