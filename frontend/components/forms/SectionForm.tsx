"use client";

import { useState, useEffect } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { api } from "@/lib/api";
import { toast } from "sonner";
import { AcademicYear, GradeLevel } from "@/types";

const sectionSchema = z.object({
  name: z.string().min(1, "Le nom est requis"),
  gradeLevelId: z.string().uuid("Niveau requis"),
  academicYearId: z.string().uuid("Année académique requise"),
  capacity: z.coerce.number().min(1, "Capacité minimale 1"),
  homeroomTeacherId: z.string().optional(),
});

type SectionFormData = z.infer<typeof sectionSchema>;

interface SectionFormProps {
  onSuccess: () => void;
}

export function SectionForm({ onSuccess }: SectionFormProps) {
  const [loading, setLoading] = useState(false);
  const [gradeLevels, setGradeLevels] = useState<GradeLevel[]>([]);
  const [academicYears, setAcademicYears] = useState<AcademicYear[]>([]);
  const { register, handleSubmit, setValue, watch, formState: { errors } } = useForm<SectionFormData>({
    resolver: zodResolver(sectionSchema),
  });

  useEffect(() => {
    const fetchData = async () => {
      const [grades, years] = await Promise.all([
        api.gradeLevels.list(),
        api.academicYears.list(),
      ]);
      setGradeLevels(grades);
      setAcademicYears(years);
      if (years.length > 0) {
        const defaultYear = years.find((y: AcademicYear) => y.isCurrent) || years[0];
        setValue("academicYearId", defaultYear.id);
      }
    };
    fetchData();
  }, [setValue]);

  const onSubmit = async (data: SectionFormData) => {
    setLoading(true);
    try {
      await api.sections.create(data);
      toast.success("Section créée avec succès");
      onSuccess();
    } catch (error: any) {
      toast.error(error.message || "Erreur lors de la création");
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
      <div className="space-y-2">
        <Label htmlFor="name">Nom de la section</Label>
        <Input id="name" placeholder="A, B, C..." {...register("name")} />
        {errors.name && <p className="text-sm text-red-500">{errors.name.message}</p>}
      </div>
      <div className="space-y-2">
        <Label htmlFor="gradeLevelId">Niveau</Label>
        <Select onValueChange={(value) => setValue("gradeLevelId", value)}>
          <SelectTrigger>
            <SelectValue placeholder="Sélectionner un niveau" />
          </SelectTrigger>
          <SelectContent>
            {gradeLevels.map((gl) => (
              <SelectItem key={gl.id} value={gl.id}>{gl.name}</SelectItem>
            ))}
          </SelectContent>
        </Select>
        {errors.gradeLevelId && <p className="text-sm text-red-500">{errors.gradeLevelId.message}</p>}
      </div>
      <div className="space-y-2">
        <Label htmlFor="academicYearId">Année académique</Label>
        <Select onValueChange={(value) => setValue("academicYearId", value)}>
          <SelectTrigger>
            <SelectValue placeholder="Sélectionner une année" />
          </SelectTrigger>
          <SelectContent>
            {academicYears.map((ay) => (
              <SelectItem key={ay.id} value={ay.id}>{ay.name}</SelectItem>
            ))}
          </SelectContent>
        </Select>
        {errors.academicYearId && <p className="text-sm text-red-500">{errors.academicYearId.message}</p>}
      </div>
      <div className="space-y-2">
        <Label htmlFor="capacity">Capacité</Label>
        <Input id="capacity" type="number" {...register("capacity")} />
        {errors.capacity && <p className="text-sm text-red-500">{errors.capacity.message}</p>}
      </div>
      <Button type="submit" className="w-full" disabled={loading}>
        {loading ? "Création..." : "Créer la section"}
      </Button>
    </form>
  );
}