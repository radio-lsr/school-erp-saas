import { auth } from "./auth";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

async function fetchAPI(endpoint: string, options: RequestInit = {}) {
  const token = auth.getToken();
  const headers: HeadersInit = {
    "Content-Type": "application/json",
    ...options.headers,
  };
  if (token) headers.Authorization = `Bearer ${token}`;
  const res = await fetch(`${API_URL}${endpoint}`, { ...options, headers });
  if (!res.ok) throw new Error(await res.text());
  if (res.status === 204) return null;
  return res.json();
}

export const api = {
  auth: {
    login: (email: string, password: string) =>
      fetchAPI("/api/auth/login", { method: "POST", body: JSON.stringify({ email, password }) }),
    register: (data: any) =>
      fetchAPI("/api/auth/register", { method: "POST", body: JSON.stringify(data) }),
  },
  students: {
    list: () => fetchAPI("/api/students"),
    get: (id: string) => fetchAPI(`/api/students/${id}`),
    create: (data: any) => fetchAPI("/api/students", { method: "POST", body: JSON.stringify(data) }),
    update: (id: string, data: any) =>
      fetchAPI(`/api/students/${id}`, { method: "PUT", body: JSON.stringify(data) }),
    delete: (id: string) => fetchAPI(`/api/students/${id}`, { method: "DELETE" }),
  },
  sections: {
    list: () => fetchAPI("/api/sections"),
    get: (id: string) => fetchAPI(`/api/sections/${id}`),
    create: (data: any) => fetchAPI("/api/sections", { method: "POST", body: JSON.stringify(data) }),
    update: (id: string, data: any) =>
      fetchAPI(`/api/sections/${id}`, { method: "PUT", body: JSON.stringify(data) }),
    delete: (id: string) => fetchAPI(`/api/sections/${id}`, { method: "DELETE" }),
  },
  gradeLevels: {
    list: () => fetchAPI("/api/grade-levels"),
  },
  academicYears: {
    list: () => fetchAPI("/api/academic-years"),
  },
  invoices: {
    list: (params?: any) => {
      const query = new URLSearchParams(params).toString();
      return fetchAPI(`/api/invoices${query ? `?${query}` : ""}`);
    },
    generate: (academicYearId?: string) =>
      fetchAPI("/api/invoices/generate", {
        method: "POST",
        body: JSON.stringify({ academic_year_id: academicYearId }),
      }),
    pay: (invoiceId: string, data: any) =>
      fetchAPI(`/api/invoices/${invoiceId}/pay`, { method: "POST", body: JSON.stringify(data) }),
  },
};