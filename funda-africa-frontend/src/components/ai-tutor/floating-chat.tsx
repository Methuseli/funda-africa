"use client"

import React, { useState, useRef, useEffect } from 'react'
import { MessageCircle, X, Send, Bot, User } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { ScrollArea } from '@/components/ui/scroll-area'
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar'
import { useAppStore } from '@/lib/store'
import { cn } from '@/lib/utils'
import 'katex/dist/katex.min.css'
import { InlineMath, BlockMath } from 'react-katex'

interface Message {
  id: string
  content: string
  sender: 'user' | 'ai'
  timestamp: string
}

const renderMessageContent = (content: string) => {
  // Simple LaTeX detection and rendering
  const parts = content.split(/(\$\$[^$]+\$\$|\$[^$]+\$)/g)
  
  return parts.map((part, index) => {
    if (part.startsWith('$$') && part.endsWith('$$')) {
      // Block math
      const latex = part.slice(2, -2)
      return <BlockMath key={index} math={latex} />
    } else if (part.startsWith('$') && part.endsWith('$')) {
      // Inline math
      const latex = part.slice(1, -1)
      return <InlineMath key={index} math={latex} />
    } else {
      return <span key={index}>{part}</span>
    }
  })
}

export default function FloatingChat() {
  const { isChatOpen, toggleChat, chatHistory, addChatMessage, currentSubtopic } = useAppStore()
  const [messages, setMessages] = useState<Message[]>([
    {
      id: '1',
      content: 'Hi! I\'m your AI tutor. I can help you understand concepts, solve problems, and answer questions about your current lesson. What would you like to know?',
      sender: 'ai',
      timestamp: new Date().toISOString()
    }
  ])
  const [inputValue, setInputValue] = useState('')
  const [isTyping, setIsTyping] = useState(false)
  const messagesEndRef = useRef<HTMLDivElement>(null)

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' })
  }

  useEffect(() => {
    scrollToBottom()
  }, [messages])

  const simulateAIResponse = (userMessage: string): string => {
    const responses = [
      `Great question! Let me help you understand that concept better. ${userMessage.includes('vector') ? 'Vectors are fundamental in linear algebra - they represent both magnitude and direction. For example, velocity is a vector because it has both speed (magnitude) and direction.' : ''}`,
      `That's an interesting point! ${userMessage.includes('limit') ? 'Limits help us understand function behavior. The limit $\\lim_{x \\to a} f(x) = L$ means f(x) approaches L as x approaches a.' : 'Let me break this down step by step for you.'}`,
      `I can see you're working on this topic! ${userMessage.includes('sort') ? 'Sorting algorithms are essential in computer science. Bubble sort has $O(n^2)$ complexity, while merge sort is $O(n \\log n)$.' : 'Would you like me to provide some examples or practice problems?'}`,
      'Excellent question! This connects to several important concepts we\'ve covered. Let me explain the relationship between these ideas.',
      `Perfect! You're asking about one of the key concepts in this area. ${userMessage.includes('Newton') ? 'Newton\'s laws form the foundation of classical mechanics. The second law $F = ma$ is particularly important for problem-solving.' : 'This is fundamental to understanding the broader topic.'}`
    ]
    
    return responses[Math.floor(Math.random() * responses.length)]
  }

  const handleSendMessage = () => {
    if (!inputValue.trim()) return

    const userMessage: Message = {
      id: Date.now().toString(),
      content: inputValue,
      sender: 'user',
      timestamp: new Date().toISOString()
    }

    setMessages(prev => [...prev, userMessage])
    setInputValue('')
    setIsTyping(true)

    // Simulate AI typing delay
    setTimeout(() => {
      const aiResponse: Message = {
        id: (Date.now() + 1).toString(),
        content: simulateAIResponse(inputValue),
        sender: 'ai',
        timestamp: new Date().toISOString()
      }
      
      setMessages(prev => [...prev, aiResponse])
      setIsTyping(false)
    }, 1000 + Math.random() * 2000)
  }

  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault()
      handleSendMessage()
    }
  }

  if (!isChatOpen) {
    return (
      <div className="fixed bottom-6 right-6 z-50">
        <Button
          onClick={toggleChat}
          size="lg"
          className="h-14 w-14 rounded-full bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 shadow-lg hover:shadow-xl transition-all duration-200"
        >
          <MessageCircle className="h-6 w-6" />
        </Button>
      </div>
    )
  }

  return (
    <div className="fixed bottom-6 right-6 z-50 w-96 h-[500px]">
      <Card className="w-full h-full shadow-2xl">
        <CardHeader className="pb-4">
          <div className="flex items-center justify-between">
            <div className="flex items-center space-x-2">
              <div className="h-8 w-8 rounded-full bg-gradient-to-r from-blue-600 to-purple-600 flex items-center justify-center">
                <Bot className="h-4 w-4 text-white" />
              </div>
              <CardTitle className="text-lg">AI Tutor</CardTitle>
            </div>
            <Button
              variant="ghost"
              size="sm"
              onClick={toggleChat}
              className="h-8 w-8 p-0"
            >
              <X className="h-4 w-4" />
            </Button>
          </div>
          <p className="text-sm text-gray-600">Ask me anything about your current lesson!</p>
        </CardHeader>
        
        <CardContent className="p-0 flex flex-col h-[400px]">
          <ScrollArea className="flex-1 p-4">
            <div className="space-y-4">
              {messages.map((message) => (
                <div
                  key={message.id}
                  className={cn(
                    'flex gap-3',
                    message.sender === 'user' ? 'justify-end' : 'justify-start'
                  )}
                >
                  {message.sender === 'ai' && (
                    <Avatar className="h-6 w-6 mt-1">
                      <AvatarFallback className="bg-gradient-to-r from-blue-600 to-purple-600 text-white text-xs">
                        AI
                      </AvatarFallback>
                    </Avatar>
                  )}
                  
                  <div
                    className={cn(
                      'max-w-[80%] rounded-lg px-3 py-2 text-sm',
                      message.sender === 'user'
                        ? 'bg-blue-600 text-white ml-auto'
                        : 'bg-gray-100 text-gray-900'
                    )}
                  >
                    <div className="prose prose-sm max-w-none">
                      {renderMessageContent(message.content)}
                    </div>
                  </div>
                  
                  {message.sender === 'user' && (
                    <Avatar className="h-6 w-6 mt-1">
                      <AvatarFallback className="bg-gray-600 text-white text-xs">
                        U
                      </AvatarFallback>
                    </Avatar>
                  )}
                </div>
              ))}
              
              {isTyping && (
                <div className="flex gap-3 justify-start">
                  <Avatar className="h-6 w-6 mt-1">
                    <AvatarFallback className="bg-gradient-to-r from-blue-600 to-purple-600 text-white text-xs">
                      AI
                    </AvatarFallback>
                  </Avatar>
                  <div className="bg-gray-100 rounded-lg px-3 py-2">
                    <div className="flex space-x-1">
                      <div className="h-2 w-2 bg-gray-400 rounded-full animate-bounce"></div>
                      <div className="h-2 w-2 bg-gray-400 rounded-full animate-bounce delay-100"></div>
                      <div className="h-2 w-2 bg-gray-400 rounded-full animate-bounce delay-200"></div>
                    </div>
                  </div>
                </div>
              )}
              
              <div ref={messagesEndRef} />
            </div>
          </ScrollArea>
          
          <div className="p-4 border-t">
            <div className="flex space-x-2">
              <Input
                value={inputValue}
                onChange={(e) => setInputValue(e.target.value)}
                onKeyPress={handleKeyPress}
                placeholder="Ask a question..."
                disabled={isTyping}
                className="flex-1"
              />
              <Button
                onClick={handleSendMessage}
                disabled={!inputValue.trim() || isTyping}
                size="sm"
                className="px-3"
              >
                <Send className="h-4 w-4" />
              </Button>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  )
}