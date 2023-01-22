import express, { Application, Request, Response } from 'express';
import cors from 'cors';

// constants
const app: Application = express();
const PORT = process.env.PORT || 4080

// Middleware
app.use(express.json())
app.use(cors())


app.get("/", (req: Request, res: Response) => {
    res.send("Healthy")
})

app.listen(PORT, () => {
    console.log(`Server is running on PORT http://localhost:${PORT}`)
})