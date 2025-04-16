import { useState, useEffect, useCallback } from 'react'
import { debounce } from 'lodash'
import './index.css'

// Custom hook to check username availability
function useUsernameCheck() {
  const [username, setUsername] = useState('')
  const [available, setAvailable] = useState(null)
  const [checking, setChecking] = useState(false)

  const serverUri = "http://localhost:8080"
  const checkUsername = useCallback(
    debounce(async (name) => {
      if (!name) {
        setAvailable(null)
        setChecking(false)
        return
      }
      setChecking(true)
      try {
        const res = await fetch(`${serverUri}/check-username`, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({ username: name }),
        })
        const data = await res.json()
        setAvailable(data.available)
      } catch (err) {
        console.error('Check failed', err)
        setAvailable(null)
      } finally {
        setChecking(false)
      }
    }, 300),
    []
  )

  useEffect(() => {
    checkUsername(username)
  }, [username, checkUsername])

  return { username, setUsername, available, checking }
}

export default function App() {
  const { username, setUsername, available, checking } = useUsernameCheck()
  const [status, setStatus] = useState('')

  const handleSubmit = async (e) => {
    e.preventDefault()

    // Final check before submitting
    const resCheck = await fetch(`${serverUri}}/check-username`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username }),
    })
    const checkData = await resCheck.json()

    if (!checkData.available) {
      setStatus('❌ Username is taken')
      return
    }

    const res = await fetch(`${serverUri}/register`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username }),
    })

    if (res.ok) {
      setStatus('✅ Registered successfully')
    } else {
      const msg = await res.text()
      setStatus(`❌ ${msg}`)
    }
  }

  return (
    <div className="min-h-screen bg-gray-50 flex flex-col items-center justify-start">
      {/* Landing Section */}
      <section className="w-full bg-gradient-to-r from-indigo-500 to-blue-500 text-white py-24 text-center px-4">
        <h1 className="text-4xl md:text-5xl font-bold mb-4">BloomFilter.io 🚀</h1>
        <p className="text-lg md:text-xl mb-8">
          Check millions of usernames in milliseconds.
        </p>
        <a
          href="#signup"
          className="bg-white text-blue-600 px-6 py-3 rounded-full font-medium shadow hover:bg-gray-100 transition"
        >
          Try Now
        </a>
      </section>

      {/* Signup Form */}
      <section
        id="signup"
        className="w-full max-w-md bg-white p-8 rounded-2xl shadow-xl mt-[-60px] z-10"
      >
        <h2 className="text-2xl font-bold mb-4 text-center">Sign Up</h2>
        <form onSubmit={handleSubmit} className="space-y-4">
          <input
            type="text"
            placeholder="Enter username"
            className="w-full border px-4 py-2 rounded-xl"
            value={username}
            onChange={(e) => {
              setUsername(e.target.value)
              setStatus('')
            }}
          />

          {/* Status Text */}
          {checking && (
            <div className="flex items-center space-x-2 text-blue-500 text-sm">
              <div className="w-4 h-4 border-2 border-blue-500 border-t-transparent rounded-full animate-spin"></div>
              <span>Checking availability...</span>
            </div>
          )}
          {!checking && available !== null && (
            <p className={`text-sm ${available ? 'text-green-600' : 'text-red-600'}`}>
              {available ? '✅ Username is available' : '❌ Username is taken'}
            </p>
          )}

          <button
            type="submit"
            className={`bg-blue-600 text-white px-4 py-2 rounded-xl w-full transition ${
              (!username || available === false || checking)
                ? 'opacity-50 cursor-not-allowed'
                : 'hover:bg-blue-700'
            }`}
            disabled={!username || available === false || checking}
          >
            {checking ? 'Checking...' : 'Register'}
          </button>

          {status && <p className="text-sm text-center">{status}</p>}
        </form>
      </section>

      {/* Footer */}
      <footer className="mt-10 text-gray-500 text-sm">
        Built with ❤️ using Go + React + Bloom Filter
      </footer>
    </div>
  )
}
