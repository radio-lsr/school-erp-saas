export interface Tenant {
  id: string;
  name: string;
  subdomain: string;
  createdAt: string;
  updatedAt: string;
}

export interface User {
  id: string;
  email: string;
  fullName: string;
  role: "admin" | "teacher" | "parent" | "student";
  tenantId: string;
}

export interface Student {
  id: string;
  tenantId: string;
  userId?: string;
  firstName: string;
  lastName: string;
  birthDate?: string;
  gender: string;
  createdAt: string;
  updatedAt: string;
}

export interface GradeLevel {
  id: string;
  tenantId: string;
  name: string;
  cycle: "maternelle" | "primaire" | "secondaire";
  displayOrder: number;
}

export interface AcademicYear {
  id: string;
  tenantId: string;
  name: string;
  startDate: string;
  endDate: string;
  isCurrent: boolean;
}

export interface Section {
  id: string;
  tenantId: string;
  gradeLevelId: string;
  academicYearId: string;
  name: string;
  capacity: number;
  homeroomTeacherId?: string;
  gradeLevelName?: string;
  homeroomTeacherName?: string;
  createdAt: string;
  updatedAt: string;
}

export interface Enrollment {
  id: string;
  tenantId: string;
  studentId: string;
  sectionId: string;
  enrollmentDate: string;
  status: string;
}

export interface FeeStructure {
  id: string;
  tenantId: string;
  gradeLevelId: string;
  academicYearId: string;
  name: string;
  totalAmount: number;
  currency: "CDF" | "USD";
}

export interface FeeInstallment {
  id: string;
  feeStructureId: string;
  periodName: string;
  amount: number;
  currency: "CDF" | "USD";
  dueDate: string;
}

export interface Invoice {
  id: string;
  invoiceNumber: string;
  studentId: string;
  student?: Student;
  feeInstallmentId: string;
  totalAmount: number;
  currency: "CDF" | "USD";
  status: "draft" | "sent" | "partially_paid" | "paid" | "overdue";
  issuedDate: string;
  dueDate: string;
  createdAt: string;
  updatedAt: string;
}

export interface Payment {
  id: string;
  invoiceId: string;
  amountPaid: number;
  currencyPaid: "CDF" | "USD";
  paymentDate: string;
  paymentMethod: "cash" | "mobile_money" | "bank_transfer";
  reference?: string;
  exchangeRate?: number;
}