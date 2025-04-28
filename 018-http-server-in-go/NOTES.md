## Takeaways

- Go servers are great for performance whether the workload is I/O or CPU-bound
- Node.js and Express work well for I/O-bound tasks, but struggle with CPU-bound tasks
- Python and Django/Flask do just fine with I/O bound tasks, but frankly, no one picks Python for its performance

## Stateful Handlers

It's frequently useful to have a way to store and access state in our handlers. For example, we might want to keep track of the number of requests we've received, or we may want to pass around an open connection to a database, or credentials to an API.

## Middleware

Middleware is a way to wrap a handler with additional functionality. It is a common pattern in web applications that allows us to write DRY code.
For example, we can write a middleware that logs every request to the server. We can then wrap our handler with this middleware and every request will be logged without us having to write the logging code in every handler.

Here are examples of the middleware that we've written so far.

```go
// Keeping Track of the Number of Times a Handler Has Been Called
func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

// Logging Every Request
// We haven't written this one yet, but it would look something like this:

func middlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
```

## Rules and Definitions
### Fixed URL Paths
A pattern that exactly matches the URL path. For example, if you have a pattern /about, it will match the URL path /about and no other paths.

### Subtree Paths
If a pattern ends with a slash /, it matches all URL paths that have the same prefix. For example, a pattern /images/ matches /images/, /images/logo.png, and /images/css/style.css. As we saw with our /app/ path, this is useful for serving a directory of static files or for structuring your application into sub-sections.

### Longest Match Wins
If more than one pattern matches a request path, the longest match is chosen. This allows more specific handlers to override more general ones. For example, if you have patterns / (root) and /images/, and the request path is /images/logo.png, the /images/ handler will be used because it's the longest match.

### Host-Specific Patterns
We won't be using this but be aware that patterns can also start with a hostname (e.g., www.example.com/). This allows you to serve different content based on the Host header of the request. If both host-specific and non-host-specific patterns match, the host-specific pattern takes precedence.

## There is always a trade-off.

### Pros for Monoliths
- Simpler to get started with
- Easier to deploy new versions because everything is always in sync
- In the case of the data being embedded in the HTML, the performance can result in better UX and SEO

### Pros for Decoupled Architectures
- Easier to scale as traffic grows
- Easier to practice good separation of concerns as the codebase grows
- Can be hosted on separate servers and using separate technologies
- Embedding data in the HTML is still possible with pre-rendering (similar to how Next.js works), it's just more complicated

## The Context Package
The context package is a part of Go's standard library. It does several things, but the most important thing is that it handles timeouts. All of SQLC's database queries accept a context.Context as their first argument:

```go
user, err := cfg.db.CreateUser(r.Context(), params.Email)
```
By passing your handler's http.Request.Context() to the query, the library will automatically cancel the database query if the HTTP request is canceled or times out.

## Authentication
Authentication is the process of verifying who a user is. If you don't have a secure authentication system, your back-end systems will be open to attack!

### Passwords Should Be Strong
The most important factor for the strength of a password is its entropy. Entropy is a measure of how many possible combinations of characters there are in a string. To put it simply:
- The longer the password the better
- Special characters and capitals should always be allowed
- Special characters and capitals aren't as important as length

### Passwords Should Never Be Stored in Plain Text
The most critical thing we can do to protect our users' passwords is to never store them in plain text. We should use cryptographically strong key derivation functions (which are a special class of hash functions) to store passwords in a way that prevents them from being read by anyone who gets access to your database.

Bcrypt is a great choice. SHA-256 and MD5 are not.

## Types of Authentication
Here are a few of the most common authentication methods you'll see in the wild:

1. Password + ID (username, email, etc.)
2. 3rd Party Authentication ("Sign in with Google", "Sign in with GitHub", etc)
3. Magic Links
4. API Keys

## You can generate a nice long random string on the command line like this:
```bash
openssl rand -base64 64
```

## Revoking JWTs
One of the main benefits of JWTs is that they're stateless. The server doesn't need to keep track of which users are logged in via JWT. The server just needs to issue a JWT to a user and the user can use that JWT to authenticate themselves. Statelessness is fast and scalable because your server doesn't need to consult a database to see if a user is currently logged in.

However, that same benefit poses a potential problem. JWTs can't be revoked. If a user's JWT is stolen, there's no easy way to stop the JWT from being used. JWTs are just a signed string of text.

The JWTs we've been using so far are more specifically access tokens. Access tokens are used to authenticate a user to a server, and they provide access to protected resources. Access tokens are:
- Stateless
- Short-lived (15m-24h)
- Irrevocable

They must be short-lived because they can't be revoked. The shorter the lifespan, the more secure they are. Trouble is, this can create a poor user experience. We don't want users to have to log in every 15 minutes.

### A Solution: Refresh Tokens
Refresh tokens don't provide access to resources directly, but they can be used to get new access tokens. Refresh tokens are much longer lived, and importantly, they can be revoked. They are:
- Stateful
- Long-lived (24h-60d)
- Revocable
Now we get the best of both worlds! Our endpoints and servers that provide access to protected resources can use access tokens, which are fast, stateless, simple, and scalable. On the other hand, refresh tokens are used to keep users logged in for longer periods of time, and they can be revoked if a user's access token is compromised.

## Cookies
HTTP cookies are one of the most talked about, but least understood, aspects of the web.

When cookies are talked about in the news, they're usually implied to simply be privacy-stealing bad guys. While cookies can certainly invade your privacy, that's not what they are.

### What Is an HTTP Cookie?
A cookie is a small piece of data that a server sends to a client. The client then dutifully stores the cookie and sends it back to the server on subsequent requests.

Cookies can store any arbitrary data:

- A user's name or other tracking information
- A JWT (refresh and access tokens)
- Items in a shopping cart
- etc.
The server decides what to put in a cookie, and the client's job is simply to store it and send it back.

### How Do Cookies Work?
Simply put, cookies work through HTTP headers.

Cookies are sent from the server to the client in the Set-Cookie header. Cookies are most popular for web (browser-based) applications because browsers automatically send any cookies they have back to the server in the Cookie header.

### Why Aren't We Using Cookies?
Simply put, Sometime the API is designed to be consumed by mobile apps and other servers. Cookies are primarily for browsers.

A good use-case for cookies is to serve as a more strict and secure transport layer for JWTs within the context of a browser-based application.

For example, when using httpOnly.cookies, you can ensure that 3rd party JavaScript that's being executed on your website can't access any cookies. That's a lot better than storing JWTs in the browser's local storage, where it's easily accessible to any JavaScript running on the page.

## Authorization
While authentication is about verifying who a user is, authorization is about verifying what a user is allowed to do.

For example, a hypothetical YouTuber ThePrimeagen should be allowed to edit and delete the videos on his account, and everyone should be allowed to view them. Another absolutely-not-real YouTuber TEEJ should be able to view ThePrimeagen's videos, but not edit or delete them.

Authorization logic is just the code that enforces these kinds of rules.

Authentication vs Authorization
- Verify who a user is, typically by asking for a password, api key, or other credentials.
- Only allow a verified user to perform actions that they are allowed to perform. Sometimes it's based on exactly who they are, but often it's based on a role, like "admin" or "owner".

## Webhooks
A webhook is just an event that's sent to your server by an external service when something happens.

For example, here at Boot.dev we use Stripe as a third-party payment processor. When a student makes a payment, Stripe sends a webhook to the Boot.dev servers so that we can unlock the student's membership.

- Student makes a payment to stripe
- Stripe processes the payment
- If the payment is successful, Stripe sends an HTTP POST request to https://api.boot.dev/stripe/webhook (that's not the real URL, but you get the idea)

That's it! The only real difference between a webhook and a typical HTTP request is that the system making the request is an automated system, not a human loading a webpage or web app. As such, webhook handlers must be idempotent because the system on the other side may retry the request multiple times.

### Idempo... What?
Idempotent, or "idempotence", is a fancy word that means "the same result no matter how many times you do it". For example, your typical POST /api/chirps (create a chirp) endpoint will not be idempotent. If you send the same request twice, you'll end up with two chirps with the same information but different IDs.

Webhooks, on the other hand, should be idempotent, and it's typically easy to build them this way because the client sends some kind of "event" and usually provides its own unique ID.

A webhook is just an event that's sent to your server by an external service. There are just a couple of things to keep in mind when building a webhook handler:

- The third-party system will probably retry requests multiple times, so your handler should be idempotent.
- Be extra careful to never "acknowledge" a webhook request unless you processed it successfully. By sending a 2XX code, you're telling the third-party system that you processed the request successfully, and they'll stop retrying it.
- When you're writing a server, you typically get to define the API. However, when you're integrating a webhook from a service like Stripe, you'll probably need to adhere to their API: they'll tell you what shape the events will be sent in.

### Are Webhooks and Websockets the Same Thing?
Nope!
- A websocket is a persistent connection between a client and a server.
- Websockets are typically used for real-time communication, like chat apps. Webhooks are a one-way communication from a third-party service to your server.