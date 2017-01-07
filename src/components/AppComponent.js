import Inferno from 'inferno';

export function App({ children }) {
    return (
        <div>
            <h1>Password Checker</h1>
            <p>Enter your username and password to check if it&apos;s correct.</p>
            { children }
        </div>
    );
}
