const server = Bun.serve({
    port: 3000,
    fetch(req: Request, server: Server): Response {
        return new Response('Hello World!');
    }
});