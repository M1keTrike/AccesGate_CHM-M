import { User } from './IUsers';
import { NfcCard } from './iNfc_cards';

export interface NFCCardAssignment {
  id: number;
  user_id: User['id'];
  card_uid: NfcCard['card_uid'];
  assigned_at: Date;
  is_active: boolean;
}