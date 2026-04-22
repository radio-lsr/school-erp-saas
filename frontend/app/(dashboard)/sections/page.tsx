"use client";

import { useEffect, useState } from "react";
import { Button } from "@/components/ui/button";
import { Plus } from "lucide-react";
import { DataTable } from "@/components/tables/DataTable";
import { SectionForm } from "@/components/forms/SectionForm";
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog";
import { api } from "@/lib/api";
import { Section } from "@/types";
import { ColumnDef } from "@tanstack/react-table";

const columns: ColumnDef<Section>[] = [
  { accessorKey: "name", header: "Section" },
  { accessorKey: "gradeLevelName", header: "Niveau" },
  { accessorKey: "capacity", header: "Capacité" },
  {
    accessorKey: "homeroomTeacherName",
    header: "Titulaire",
    cell: ({ row }) => row.original.homeroomTeacherName || "-",
  },
];

export default function SectionsPage() {
  const [sections, setSections] = useState<Section[]>([]);
  const [loading, setLoading] = useState(true);
  const [open, setOpen] = useState(false);

  const fetchSections = async () => {
    setLoading(true);
    try {
      const data = await api.sections.list();
      setSections(data);
    } catch (error) {
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchSections();
  }, []);

  const handleSectionCreated = () => {
    setOpen(false);
    fetchSections();
  };

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold">Sections / Classes</h1>
          <p className="text-muted-foreground">Gérez les sections (ex: 6ème A)</p>
        </div>
        <Dialog open={open} onOpenChange={setOpen}>
          <DialogTrigger asChild>
            <Button>
              <Plus className="mr-2 h-4 w-4" />
              Nouvelle section
            </Button>
          </DialogTrigger>
          <DialogContent className="sm:max-w-[500px]">
            <DialogHeader>
              <DialogTitle>Créer une section</DialogTitle>
            </DialogHeader>
            <SectionForm onSuccess={handleSectionCreated} />
          </DialogContent>
        </Dialog>
      </div>

      <DataTable
        columns={columns}
        data={sections}
        loading={loading}
        searchPlaceholder="Rechercher une section..."
        searchColumn="name"
      />
    </div>
  );
}