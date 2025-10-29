export interface UsersInterface { // เครื่องหมาย ? หลังชื่อ field เช่น ID? หมายถึง ไม่จำเป็นต้องมี field นี้ก็ได้ (เรียกว่า optional property)
  ID?: number;
  FirstName?: string;
  LastName?: string;
  Email?: string;
  Phone?: string;
  Age?: number;
  BirthDay?: string;
  GenderID?: number;
  Address?: string;
  Password?: string;
}