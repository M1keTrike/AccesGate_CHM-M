export {};

declare global {
  interface Navigator {
    bluetooth: {
      requestDevice(options?: RequestDeviceOptions): Promise<BluetoothDevice>;
    };
  }

  interface BluetoothDevice {
    name?: string;
    gatt?: BluetoothRemoteGATTServer;
  }

  interface BluetoothRemoteGATTServer {
    connect(): Promise<BluetoothRemoteGATTServer>;
    getPrimaryService(
      serviceUUID: BluetoothServiceUUID
    ): Promise<BluetoothRemoteGATTService>;
    
    readonly connected: boolean; 
  }
  
  interface BluetoothRemoteGATTService {
    getCharacteristic(
      characteristicUUID: BluetoothCharacteristicUUID
    ): Promise<BluetoothRemoteGATTCharacteristic>;
  }

  interface BluetoothRemoteGATTCharacteristic {
    writeValue(value: BufferSource): Promise<void>;
    readValue(): Promise<DataView>;
  }

  // Añadimos tipos específicos para UUID de Bluetooth
  type BluetoothServiceUUID = string | number | UUID;
  type BluetoothCharacteristicUUID = string | number | UUID;
  type UUID = string;
}
