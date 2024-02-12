import Register from "@/components/auth/Register";
import Sigin from "@/components/auth/Sigin";
import { Tabs, TabsList, TabsTrigger, TabsContent } from "@/components/ui/tabs";

export default function AuthPage() {
  return (
    <div className="w-full h-full flex flex-col items-center justify-center gap-4">
      <Tabs defaultValue="register" className="w-[400px]">
        <TabsList>
          <TabsTrigger value="register">Register</TabsTrigger>
          <TabsTrigger value="signin">Sign in</TabsTrigger>
        </TabsList>
        <TabsContent value="signin">
          <Sigin />
        </TabsContent>
        <TabsContent value="register">
          <Register />
        </TabsContent>
      </Tabs>
    </div>
  );
}
