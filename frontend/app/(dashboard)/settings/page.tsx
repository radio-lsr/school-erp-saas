"use client";

import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Button } from "@/components/ui/button";

export default function SettingsPage() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-3xl font-bold">Paramètres</h1>
        <p className="text-muted-foreground">Configuration de votre établissement</p>
      </div>

      <div className="grid gap-6">
        <Card>
          <CardHeader>
            <CardTitle>Informations de l'école</CardTitle>
            <CardDescription>Modifiez les détails de votre établissement</CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="space-y-2">
              <Label htmlFor="schoolName">Nom de l'école</Label>
              <Input id="schoolName" defaultValue="Collège Saint Joseph" />
            </div>
            <div className="space-y-2">
              <Label htmlFor="address">Adresse</Label>
              <Input id="address" defaultValue="123, Avenue de la Paix, Kinshasa" />
            </div>
            <div className="space-y-2">
              <Label htmlFor="phone">Téléphone</Label>
              <Input id="phone" defaultValue="+243 123 456 789" />
            </div>
            <Button>Sauvegarder</Button>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Taux de change</CardTitle>
            <CardDescription>Configurez le taux USD/CDF pour les paiements</CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <div className="space-y-2">
              <Label htmlFor="rate">1 USD = ? CDF</Label>
              <Input id="rate" type="number" defaultValue="2800" />
            </div>
            <Button>Mettre à jour</Button>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}