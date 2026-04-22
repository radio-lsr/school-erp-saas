"use client";

import { useState } from "react";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "@/components/ui/select";
import { api } from "@/lib/api";
import { toast } from "sonner";

const studentSchema = z.object({
  firstName: z.string().min(1, "Le prénom est requis"),
  lastName: z.string().min(1, "Le nom est requis"),
  birthDate: z.string().optional(),
  gender: z.enum(["M", "F"]),
});

type StudentFormData = z.infer<typeof studentSchema>;

interface StudentFormProps {
  onSuccess: () => void;
}

export function StudentForm({ onSuccess }: StudentFormProps) {
  const [loading, setLoading] = useState(false);
  const { register, handleSubmit, setValue, watch, formState: { errors } } = useForm<StudentFormData>({
    resolver: zodResolver(studentSchema),
    defaultValues: { gender: "M" },
  });

  const gender = watch("gender");

  const onSubmit = async (data: StudentFormData) => {
    setLoading(true);
    try {
      await api.students.create(data);
      toast.success("Élève ajouté avec succès");
      onSuccess();
    } catch (error: any) {
      toast.error(error.message || "Erreur lors de l'ajout");
    } finally {
      setLoading(false);
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
      <div className="space-y-2">
        <Label htmlFor="firstName">Prénom</Label>
        <Input id="firstName" {...register("firstName")} />
        {errors.firstName && <p className="text-sm text-red-500">{errors.firstName.message}</p>}
      </div>
      <div className="space-y-2">
        <Label htmlFor="lastName">Nom</Label>
        <Input id="lastName" {...register("lastName")} />
        {errors.lastName && <p className="text-sm text-red-500">{errors.lastName.message}</p>}
      </div>
      <div className="space-y-2">
        <Label htmlFor="birthDate">Date de naissance</Label>
        <Input id="birthDate" type="date" {...register("birthDate")} />
      </div>
      <div className="space-y-2">
        <Label htmlFor="gender">Genre</Label>
        <Select value={gender} onValueChange={(value) => setValue("gender", value as "M" | "F")}>
          <SelectTrigger>
            <SelectValue />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="M">Masculin</SelectItem>
            <SelectItem value="F">Féminin</SelectItem>
          </SelectContent>
        </Select>
      </div>
      <Button type="submit" className="w-full" disabled={loading}>
        {loading ? "Ajout en cours..." : "Ajouter l'élève"}
      </Button>
    </form>
  );
}