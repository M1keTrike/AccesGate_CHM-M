export interface EventAttendee {
    id: number;
    user_id: number;
    event_id: number;
    registered_at: Date;
    attended: boolean;
    nfc_id: string; // Add this field
    user?: {
        id: number;
        name: string;
        email: string;
    };
}

export interface Event {
    id: number;
    name: string;
    description: string;
    start_time: Date;
    end_time: Date;
    created_by: number;
    created_at: Date;
}