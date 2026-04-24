"use client";

import { useEffect, useState } from "react";
import { StatsCards } from "@/components/dashboard/StatsCards";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { api } from "@/lib/api";
import { Student, Invoice } from "@/types";

export default function DashboardPage() {
  const [stats, setStats] = useState({
    totalStudents: 0,
    totalSections: 0,
    pendingInvoices: 0,
    revenue: 0,
  });
  const [recentStudents, setRecentStudents] = useState<Student[]>([]);
  const [recentInvoices, setRecentInvoices] = useState<Invoice[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchDashboardData = async () => {
      try {
        const [studentsRes, sectionsRes, invoicesRes] = await Promise.all([
          api.students.list(),
          api.sections.list(),
          api.invoices.list({ status: "pending" }),
        ]);
        // Calcul des stats...
        setStats({
          totalStudents: studentsRes.length,
          totalSections: sectionsRes.length,
          pendingInvoices: invoicesRes.length,
          revenue: 0, // À calculer
        });
        setRecentStudents(studentsRes.slice(0, 5));
        setRecentInvoices(invoicesRes.slice(0, 5));
      } catch (error) {
        console.error(error);
      } finally {
        setLoading(false);
      }
    };
    fetchDashboardData();
  }, []);

  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-3xl font-bold">Tableau de bord</h1>
        <p className="text-muted-foreground">
          Bienvenue dans votre espace de gestion scolaire
        </p>
      </div>

      <StatsCards stats={stats} loading={loading} />

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <Card>
          <CardHeader>
            <CardTitle>Élèves récents</CardTitle>
            <CardDescription>Derniers élèves inscrits</CardDescription>
          </CardHeader>
          <CardContent>
            {loading ? (
              <p>Chargement...</p>
            ) : recentStudents.length === 0 ? (
              <p className="text-muted-foreground">Aucun élève</p>
            ) : (
              <ul className="space-y-2">
                {recentStudents.map((student) => (
                  <li key={student.id} className="flex justify-between">
                    <span>{student.firstName} {student.lastName}</span>
                    <span className="text-sm text-muted-foreground">
                      {new Date(student.createdAt).toLocaleDateString()}
                    </span>
                  </li>
                ))}
              </ul>
            )}
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Factures en attente</CardTitle>
            <CardDescription>Paiements à recevoir</CardDescription>
          </CardHeader>
          <CardContent>
            {loading ? (
              <p>Chargement...</p>
            ) : recentInvoices.length === 0 ? (
              <p className="text-muted-foreground">Aucune facture en attente</p>
            ) : (
              <ul className="space-y-2">
                {recentInvoices.map((invoice) => (
                  <li key={invoice.id} className="flex justify-between">
                    <span>Facture #{invoice.invoiceNumber}</span>
                    <span className="font-medium">
                      {invoice.totalAmount} {invoice.currency}
                    </span>
                  </li>
                ))}
              </ul>
            )}
          </CardContent>
        </Card>
      </div>
    </div>
  );
}