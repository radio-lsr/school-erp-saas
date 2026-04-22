"use client";

import { useEffect, useState } from "react";
import { Button } from "@/components/ui/button";
import { Plus } from "lucide-react";
import { DataTable } from "@/components/tables/DataTable";
import { StudentForm } from "@/components/forms/StudentForm";
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from "@/components/ui/dialog";
import { api } from "@/lib/api";
import { Student } from "@/types";
import { ColumnDef } from "@tanstack/react-table";

const columns: ColumnDef<Student>[] = [
  { accessorKey: "firstName", header: "Prénom" },
  { accessorKey: "lastName", header: "Nom" },
  { accessorKey: "gender", header: "Genre" },
  {
    accessorKey: "birthDate",
    header: "Date de naissance",
    cell: ({ row }) => row.original.birthDate ? new Date(row.original.birthDate).toLocaleDateString() : "-",
  },
];

export default function StudentsPage() {
  const [students, setStudents] = useState<Student[]>([]);
  const [loading, setLoading] = useState(true);
  const [open, setOpen] = useState(false);

  const fetchStudents = async () => {
    setLoading(true);
    try {
      const data = await api.students.list();
      setStudents(data);
    } catch (error) {
      console.error(error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchStudents();
  }, []);

  const handleStudentCreated = () => {
    setOpen(false);
    fetchStudents();
  };

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-bold">Élèves</h1>
          <p className="text-muted-foreground">Gérez la liste des élèves</p>
        </div>
        <Dialog open={open} onOpenChange={setOpen}>
          <DialogTrigger asChild>
            <Button>
              <Plus className="mr-2 h-4 w-4" />
              Nouvel élève
            </Button>
          </DialogTrigger>
          <DialogContent className="sm:max-w-[500px]">
            <DialogHeader>
              <DialogTitle>Ajouter un élève</DialogTitle>
            </DialogHeader>
            <StudentForm onSuccess={handleStudentCreated} />
          </DialogContent>
        </Dialog>
      </div>

      <DataTable
        columns={columns}
        data={students}
        loading={loading}
        searchPlaceholder="Rechercher un élève..."
        searchColumn="lastName"
      />
    </div>
  );
}