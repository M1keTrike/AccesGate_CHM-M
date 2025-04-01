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
      serviceUUID: string | number
    ): Promise<BluetoothRemoteGATTService>;
    
    readonly connected: boolean; 
  }
  

  interface BluetoothRemoteGATTService {
    getCharacteristic(
      characteristicUUID: string | number
    ): Promise<BluetoothRemoteGATTCharacteristic>;
  }

  interface BluetoothRemoteGATTCharacteristic {
    writeValue(value: BufferSource): Promise<void>;
    readValue(): Promise<DataView>;
  }
}
