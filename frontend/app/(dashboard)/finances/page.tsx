"use client";

import { useEffect, useState } from "react";
import { Button } from "@/components/ui/button";
import { DataTable } from "@/components/tables/DataTable";
import { api } from "@/lib/api";
import { Invoice } from "@/types";
import { ColumnDef } from "@tanstack/react-table";
import { Badge } from "@/components/ui/badge";
import { toast } from "sonner";

const columns: ColumnDef<Invoice>[] = [
  { accessorKey: "invoiceNumber", header: "N° Facture" },
  {
    accessorKey: "studentName",
    header: "Élève",
    cell: ({ row }) => row.original.student?.firstName + " " + row.original.student?.lastName,
  },
  {
    accessorKey: "totalAmount",
    header: "Montant",
    cell: ({ row }) => `${row.original.totalAmount} ${row.original.currency}`,
  },
  {
    accessorKey: "dueDate",
    header: "Échéance",
    cell: ({ row }) => new Date(row.original.dueDate).toLocaleDateString(),
  },
  {
    accessorKey: "status",
    header: "Statut",
    cell: ({ row }) => {
      const status = row.original.status;
      return (
        <Badge variant={status === "paid" ? "success" : status === "partially_paid" ? "warning" : "destructive"}>
          {status === "paid" ? "Payée" : status === "partially_paid" ? "Partiel" : "Impayée"}
        </Badge>
      );
    },
  },
  {
    id: "actions",
    cell: ({ row }) => (
      <Button
        variant="outline"
        size="sm"
        disabled={row.original.status === "paid"}
        onClick={() => handlePay(row.original)}
      >
        Payer
      </Button>
    ),
  },
];

const handlePay = async (invoice: Invoice) => {
  // Ouvrir une modale de paiement Mobile Money
  toast.info(`Paiement pour la facture ${invoice.invoiceNumber}`);
  // Implémentation complète avec modale à ajouter
};

export default function FinancesPage() {
  const [invoices, setInvoices] = useState<Invoice[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchInvoices();
  }, []);

  const fetchInvoices = async () => {
    setLoading(true);
    try {
      const data = await api.invoices.list();
      setInvoices(data);
    } catch (error) {
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold">Finances</h1>
          <p className="text-muted-foreground">Gérez les factures et paiements</p>
        </div>
        <Button onClick={() => api.invoices.generate()}>Générer les factures</Button>
      </div>

      <DataTable
        columns={columns}
        data={invoices}
        loading={loading}
        searchPlaceholder="Rechercher une facture..."
        searchColumn="invoiceNumber"
      />
    </div>
  );
}