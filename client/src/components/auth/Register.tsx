import { Label } from '@/components/ui/label';
import { useState, FormEvent } from 'react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';
import { register } from '@/api/authService';

export default function Register() {
    const [userName, setUserName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');
    const [message, setMessage] = useState('');

    const handleSubmit = async (event: FormEvent) => {
        event.preventDefault();
        if (password !== confirmPassword) {
            setMessage("Les mots de passe ne correspondent pas.");
            return;
        }
        try {
            const data = await register(userName, email, password);
            if (data) {
                setMessage("Register successful");
            }
        } catch (error) {
            setMessage(error instanceof Error ? error.message : "Unknown error");
        }
    };

    return (
        <form onSubmit={handleSubmit} className="flex flex-col gap-4">
            <div>
                <Label>Username</Label>
                <Input
                    value={userName}
                    onChange={(e) => setUserName(e.target.value)}
                    placeholder="Enter your username"
                    type="text"
                />
            </div>
            <div>
                <Label>Email</Label>
                <Input
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    placeholder="Enter your email"
                    type="email"
                />
            </div>
            <div>
                <Label>Password</Label>
                <Input
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    placeholder="Enter your password"
                    type="password"
                />
            </div>
            <div>
                <Label>Confirm Password</Label>
                <Input
                    value={confirmPassword}
                    onChange={(e) => setConfirmPassword(e.target.value)}
                    placeholder="Confirm your password"
                    type="password"
                />
            </div>
            <Button type="submit">Register</Button>
            {message && <div>{message}</div>}
        </form>
    );
}
