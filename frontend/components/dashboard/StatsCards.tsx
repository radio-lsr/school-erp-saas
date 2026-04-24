import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Users, BookOpen, CreditCard, DollarSign } from "lucide-react";

interface StatsCardsProps {
  stats: { totalStudents: number; totalSections: number; pendingInvoices: number; revenue: number };
  loading: boolean;
}

export function StatsCards({ stats, loading }: StatsCardsProps) {
  const cards = [
    { title: "Total élèves", value: stats.totalStudents, icon: Users, bg: "bg-violet-100 text-violet-600" },
    { title: "Sections actives", value: stats.totalSections, icon: BookOpen, bg: "bg-emerald-100 text-emerald-600" },
    { title: "Factures impayées", value: stats.pendingInvoices, icon: CreditCard, bg: "bg-amber-100 text-amber-600" },
    { title: "Revenus (USD)", value: `$${stats.revenue.toLocaleString()}`, icon: DollarSign, bg: "bg-rose-100 text-rose-600" },
  ];

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-4">
      {cards.map((card, index) => (
        <Card key={index} className="border-0">
          <CardHeader className="flex flex-row items-center justify-between pb-2">
            <CardTitle className="text-sm font-medium">{card.title}</CardTitle>
            <div className={`p-2 rounded-full ${card.bg}`}>
              <card.icon className="h-4 w-4" />
            </div>
          </CardHeader>
          <CardContent>
            <div className="text-2xl font-bold">{loading ? "..." : card.value}</div>
          </CardContent>
        </Card>
      ))}
    </div>
  );
}