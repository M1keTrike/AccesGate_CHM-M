importScripts('https://www.gstatic.com/firebasejs/10.7.1/firebase-app-compat.js');
importScripts('https://www.gstatic.com/firebasejs/10.7.1/firebase-messaging-compat.js');

firebase.initializeApp({
  apiKey: "AIzaSyCRTQ721jWEeJU3J_BnfbvLvh7M2k3Ml8s",
  authDomain: "accessgate-ch-mm.firebaseapp.com",
  projectId: "accessgate-ch-mm",
  storageBucket: "accessgate-ch-mm.firebasestorage.app",
  messagingSenderId: "573648878950",
  appId: "1:573648878950:web:dfd83f5907355d10105ca7",
  measurementId: "G-DLPCS0L5RD"
});

const messaging = firebase.messaging();

messaging.onBackgroundMessage((payload) => {
  console.log('[firebase-messaging-sw.js] Recibido mensaje en segundo plano:', payload);
  self.registration.showNotification(payload.notification.title, {
    body: payload.notification.body,
    icon: '/assets/icons/icon-192x192.png'
  });
});
