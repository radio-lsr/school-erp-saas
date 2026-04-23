const TOKEN_KEY = "school_erp_token";
const USER_KEY = "school_erp_user";

interface User {
  id: string;
  email: string;
  fullName: string;
  role: string;
  tenantId: string;
}

export const auth = {
  setToken: (token: string) => {
    if (typeof window !== "undefined") localStorage.setItem(TOKEN_KEY, token);
  },
  getToken: (): string | null => {
    if (typeof window !== "undefined") return localStorage.getItem(TOKEN_KEY);
    return null;
  },
  setUser: (user: User) => {
    if (typeof window !== "undefined") localStorage.setItem(USER_KEY, JSON.stringify(user));
  },
  getUser: (): User | null => {
    if (typeof window !== "undefined") {
      const user = localStorage.getItem(USER_KEY);
      return user ? JSON.parse(user) : null;
    }
    return null;
  },
  isAuthenticated: (): boolean => !!auth.getToken(),
  logout: () => {
    if (typeof window !== "undefined") {
      localStorage.removeItem(TOKEN_KEY);
      localStorage.removeItem(USER_KEY);
    }
  },
  login: async (email: string, password: string) => {
    const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/auth/login`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password }),
    });
    if (!res.ok) throw new Error("Invalid credentials");
    const data = await res.json();
    auth.setToken(data.token);
    auth.setUser(data.user);
    return data;
  },
  register: async (data: any) => {
    const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/auth/register`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    });
    if (!res.ok) throw new Error(await res.text());
    return res.json();
  },
};