import type { Metadata } from "next";
import "./globals.css";
import { Toaster } from "@/components/ui/sonner";

export const metadata: Metadata = {
  title: "School ERP SaaS",
  description: "Gestion scolaire complète pour la RDC",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="fr">
      <body className="font-sans"> {/* Utilise la police système */}
        {children}
        <Toaster />
      </body>
    </html>
  );
}