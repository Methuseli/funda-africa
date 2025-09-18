"use client"

import React, { useState, useRef, useEffect } from 'react'
import { 
  Video, 
  VideoOff, 
  Mic, 
  MicOff, 
  Phone, 
  PhoneOff, 
  Monitor, 
  MonitorOff,
  Hand,
  MessageSquare,
  Users,
  Settings,
  MoreVertical,
  Send,
  Volume2,
  VolumeX,
  Maximize,
  Minimize,
  Camera,
  Share
} from 'lucide-react'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Badge } from '@/components/ui/badge'
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar'
import { ScrollArea } from '@/components/ui/scroll-area'
import { Separator } from '@/components/ui/separator'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from '@/components/ui/dropdown-menu'
import { cn } from '@/lib/utils'
import Sidebar from '@/components/layout/sidebar'

interface Participant {
  id: string
  name: string
  avatar?: string
  isInstructor: boolean
  videoEnabled: boolean
  audioEnabled: boolean
  isHandRaised: boolean
  isScreenSharing: boolean
}

interface ChatMessage {
  id: string
  participantId: string
  participantName: string
  message: string
  timestamp: string
  type: 'message' | 'system'
}

export default function LiveLessonsPage() {
  const [isVideoEnabled, setIsVideoEnabled] = useState(true)
  const [isAudioEnabled, setIsAudioEnabled] = useState(true)
  const [isHandRaised, setIsHandRaised] = useState(false)
  const [isScreenSharing, setIsScreenSharing] = useState(false)
  const [isChatOpen, setIsChatOpen] = useState(true)
  const [isFullscreen, setIsFullscreen] = useState(false)
  const [chatMessage, setChatMessage] = useState('')
  const [isConnected, setIsConnected] = useState(true)
  const [currentView, setCurrentView] = useState<'grid' | 'speaker' | 'screen'>('speaker')
  
  const videoRef = useRef<HTMLVideoElement>(null)
  const chatEndRef = useRef<HTMLDivElement>(null)

  const [participants] = useState<Participant[]>([
    {
      id: '1',
      name: 'Dr. Sarah Johnson',
      avatar: 'https://images.pexels.com/photos/774909/pexels-photo-774909.jpeg?auto=compress&cs=tinysrgb&w=400',
      isInstructor: true,
      videoEnabled: true,
      audioEnabled: true,
      isHandRaised: false,
      isScreenSharing: true
    },
    {
      id: '2',
      name: 'Alex Chen',
      avatar: 'https://images.pexels.com/photos/614810/pexels-photo-614810.jpeg?auto=compress&cs=tinysrgb&w=400',
      isInstructor: false,
      videoEnabled: true,
      audioEnabled: false,
      isHandRaised: true,
      isScreenSharing: false
    },
    {
      id: '3',
      name: 'Maria Rodriguez',
      avatar: 'https://images.pexels.com/photos/415829/pexels-photo-415829.jpeg?auto=compress&cs=tinysrgb&w=400',
      isInstructor: false,
      videoEnabled: false,
      audioEnabled: true,
      isHandRaised: false,
      isScreenSharing: false
    },
    {
      id: '4',
      name: 'John Smith',
      isInstructor: false,
      videoEnabled: true,
      audioEnabled: true,
      isHandRaised: false,
      isScreenSharing: false
    }
  ])

  const [chatMessages] = useState<ChatMessage[]>([
    {
      id: '1',
      participantId: '1',
      participantName: 'Dr. Sarah Johnson',
      message: 'Welcome everyone to today\'s Linear Algebra session!',
      timestamp: '10:00 AM',
      type: 'message'
    },
    {
      id: '2',
      participantId: 'system',
      participantName: 'System',
      message: 'Alex Chen joined the session',
      timestamp: '10:01 AM',
      type: 'system'
    },
    {
      id: '3',
      participantId: '2',
      participantName: 'Alex Chen',
      message: 'Good morning! Excited for today\'s lesson on vector spaces.',
      timestamp: '10:02 AM',
      type: 'message'
    },
    {
      id: '4',
      participantId: '3',
      participantName: 'Maria Rodriguez',
      message: 'Can you please share the slides?',
      timestamp: '10:03 AM',
      type: 'message'
    }
  ])

  const instructor = participants.find(p => p.isInstructor)
  const students = participants.filter(p => !p.isInstructor)
  const raisedHands = participants.filter(p => p.isHandRaised).length

  const handleSendMessage = () => {
    if (chatMessage.trim()) {
      // In a real app, this would send the message via WebSocket/WebRTC
      setChatMessage('')
    }
  }

  const toggleScreenShare = () => {
    setIsScreenSharing(!isScreenSharing)
    if (!isScreenSharing) {
      setCurrentView('screen')
    } else {
      setCurrentView('speaker')
    }
  }

  useEffect(() => {
    chatEndRef.current?.scrollIntoView({ behavior: 'smooth' })
  }, [chatMessages])

  return (
    <div className="flex w-full h-screen bg-gray-900">
      <Sidebar />
      
      <div className="flex-1 flex flex-col">
        {/* Header */}
        <div className="bg-gray-800 border-b border-gray-700 px-6 py-4">
          <div className="flex items-center justify-between">
            <div>
              <h1 className="text-xl font-semibold text-white">Linear Algebra - Vector Spaces</h1>
              <p className="text-gray-400 text-sm">
                Live Session • {participants.length} participants • {raisedHands > 0 && `${raisedHands} hands raised`}
              </p>
            </div>
            
            <div className="flex items-center space-x-2">
              <Badge variant="secondary" className="bg-green-100 text-green-800">
                Live
              </Badge>
              <span className="text-gray-400 text-sm">10:05 AM - 11:00 AM</span>
            </div>
          </div>
        </div>

        <div className="flex-1 flex">
          {/* Main Video Area */}
          <div className="flex-1 flex flex-col bg-gray-900">
            {/* Video Grid/Speaker View */}
            <div className="flex-1 p-4">
              {currentView === 'screen' && isScreenSharing ? (
                // Screen Share View
                <div className="h-full bg-gray-800 rounded-lg flex items-center justify-center relative">
                  <div className="w-full h-full bg-gradient-to-br from-blue-900 to-purple-900 rounded-lg flex items-center justify-center">
                    <div className="text-center text-white">
                      <Monitor className="h-16 w-16 mx-auto mb-4 opacity-50" />
                      <p className="text-lg font-medium">Screen Share Active</p>
                      <p className="text-gray-300">Dr. Sarah Johnson is sharing their screen</p>
                    </div>
                  </div>
                  
                  {/* Instructor thumbnail */}
                  <div className="absolute top-4 right-4 w-32 h-24 bg-gray-700 rounded-lg overflow-hidden">
                    <div className="w-full h-full bg-gradient-to-br from-gray-600 to-gray-800 flex items-center justify-center">
                      <Avatar className="h-12 w-12">
                        <AvatarImage src={instructor?.avatar} />
                        <AvatarFallback>DJ</AvatarFallback>
                      </Avatar>
                    </div>
                  </div>
                </div>
              ) : currentView === 'speaker' ? (
                // Speaker View
                <div className="h-full flex">
                  <div className="flex-1 bg-gray-800 rounded-lg mr-4 relative overflow-hidden">
                    <div className="w-full h-full bg-gradient-to-br from-blue-900 to-purple-900 flex items-center justify-center">
                      <Avatar className="h-32 w-32">
                        <AvatarImage src={instructor?.avatar} />
                        <AvatarFallback className="text-4xl">DJ</AvatarFallback>
                      </Avatar>
                    </div>
                    
                    {/* Speaker info overlay */}
                    <div className="absolute bottom-4 left-4 bg-black/50 backdrop-blur-sm rounded-lg px-3 py-2">
                      <div className="flex items-center space-x-2">
                        <div className="flex items-center space-x-1">
                          {instructor?.audioEnabled ? (
                            <Mic className="h-4 w-4 text-green-400" />
                          ) : (
                            <MicOff className="h-4 w-4 text-red-400" />
                          )}
                          {instructor?.videoEnabled ? (
                            <Video className="h-4 w-4 text-green-400" />
                          ) : (
                            <VideoOff className="h-4 w-4 text-red-400" />
                          )}
                        </div>
                        <span className="text-white font-medium">{instructor?.name}</span>
                        <Badge variant="secondary" className="bg-blue-100 text-blue-800 text-xs">
                          Instructor
                        </Badge>
                      </div>
                    </div>
                  </div>
                  
                  {/* Participants sidebar */}
                  <div className="w-48 space-y-2">
                    {students.map((participant) => (
                      <div key={participant.id} className="bg-gray-800 rounded-lg p-2 relative">
                        <div className="aspect-video bg-gradient-to-br from-gray-600 to-gray-800 rounded flex items-center justify-center">
                          {participant.videoEnabled ? (
                            <Avatar className="h-8 w-8">
                              <AvatarImage src={participant.avatar} />
                              <AvatarFallback>{participant.name.split(' ').map(n => n[0]).join('')}</AvatarFallback>
                            </Avatar>
                          ) : (
                            <div className="text-gray-400">
                              <VideoOff className="h-6 w-6" />
                            </div>
                          )}
                        </div>
                        
                        <div className="mt-1 flex items-center justify-between">
                          <span className="text-white text-xs truncate">{participant.name}</span>
                          <div className="flex items-center space-x-1">
                            {participant.isHandRaised && (
                              <Hand className="h-3 w-3 text-yellow-400" />
                            )}
                            {!participant.audioEnabled && (
                              <MicOff className="h-3 w-3 text-red-400" />
                            )}
                          </div>
                        </div>
                      </div>
                    ))}
                  </div>
                </div>
              ) : (
                // Grid View
                <div className="h-full grid grid-cols-2 lg:grid-cols-3 gap-4">
                  {participants.map((participant) => (
                    <div key={participant.id} className="bg-gray-800 rounded-lg relative overflow-hidden">
                      <div className="w-full h-full bg-gradient-to-br from-gray-600 to-gray-800 flex items-center justify-center">
                        {participant.videoEnabled ? (
                          <Avatar className="h-16 w-16">
                            <AvatarImage src={participant.avatar} />
                            <AvatarFallback>{participant.name.split(' ').map(n => n[0]).join('')}</AvatarFallback>
                          </Avatar>
                        ) : (
                          <div className="text-gray-400">
                            <VideoOff className="h-8 w-8" />
                          </div>
                        )}
                      </div>
                      
                      <div className="absolute bottom-2 left-2 bg-black/50 backdrop-blur-sm rounded px-2 py-1">
                        <div className="flex items-center space-x-1">
                          <span className="text-white text-sm font-medium">{participant.name}</span>
                          {participant.isInstructor && (
                            <Badge variant="secondary" className="bg-blue-100 text-blue-800 text-xs">
                              Instructor
                            </Badge>
                          )}
                        </div>
                      </div>
                      
                      <div className="absolute top-2 right-2 flex items-center space-x-1">
                        {participant.isHandRaised && (
                          <div className="bg-yellow-500 rounded-full p-1">
                            <Hand className="h-3 w-3 text-white" />
                          </div>
                        )}
                        {!participant.audioEnabled && (
                          <div className="bg-red-500 rounded-full p-1">
                            <MicOff className="h-3 w-3 text-white" />
                          </div>
                        )}
                      </div>
                    </div>
                  ))}
                </div>
              )}
            </div>

            {/* Controls Bar */}
            <div className="bg-gray-800 border-t border-gray-700 px-6 py-4">
              <div className="flex items-center justify-between">
                <div className="flex items-center space-x-2">
                  {/* View Controls */}
                  <div className="flex items-center space-x-1 bg-gray-700 rounded-lg p-1">
                    <Button
                      variant={currentView === 'speaker' ? 'default' : 'ghost'}
                      size="sm"
                      onClick={() => setCurrentView('speaker')}
                      className="text-white"
                    >
                      Speaker
                    </Button>
                    <Button
                      variant={currentView === 'grid' ? 'default' : 'ghost'}
                      size="sm"
                      onClick={() => setCurrentView('grid')}
                      className="text-white"
                    >
                      Grid
                    </Button>
                  </div>
                </div>

                {/* Main Controls */}
                <div className="flex items-center space-x-2">
                  <Button
                    variant={isAudioEnabled ? 'default' : 'destructive'}
                    size="sm"
                    onClick={() => setIsAudioEnabled(!isAudioEnabled)}
                    className="w-12 h-12 rounded-full"
                  >
                    {isAudioEnabled ? <Mic className="h-5 w-5" /> : <MicOff className="h-5 w-5" />}
                  </Button>
                  
                  <Button
                    variant={isVideoEnabled ? 'default' : 'destructive'}
                    size="sm"
                    onClick={() => setIsVideoEnabled(!isVideoEnabled)}
                    className="w-12 h-12 rounded-full"
                  >
                    {isVideoEnabled ? <Video className="h-5 w-5" /> : <VideoOff className="h-5 w-5" />}
                  </Button>
                  
                  <Button
                    variant={isHandRaised ? 'default' : 'outline'}
                    size="sm"
                    onClick={() => setIsHandRaised(!isHandRaised)}
                    className={cn(
                      "w-12 h-12 rounded-full",
                      isHandRaised && "bg-yellow-500 hover:bg-yellow-600"
                    )}
                  >
                    <Hand className="h-5 w-5" />
                  </Button>
                  
                  <Button
                    variant={isScreenSharing ? 'default' : 'outline'}
                    size="sm"
                    onClick={toggleScreenShare}
                    className="w-12 h-12 rounded-full"
                  >
                    {isScreenSharing ? <MonitorOff className="h-5 w-5" /> : <Monitor className="h-5 w-5" />}
                  </Button>
                  
                  <Button
                    variant="destructive"
                    size="sm"
                    onClick={() => setIsConnected(false)}
                    className="w-12 h-12 rounded-full"
                  >
                    <PhoneOff className="h-5 w-5" />
                  </Button>
                </div>

                <div className="flex items-center space-x-2">
                  <Button
                    variant="ghost"
                    size="sm"
                    onClick={() => setIsChatOpen(!isChatOpen)}
                    className="text-white"
                  >
                    <MessageSquare className="h-5 w-5 mr-2" />
                    Chat
                  </Button>
                  
                  <Button
                    variant="ghost"
                    size="sm"
                    className="text-white"
                  >
                    <Users className="h-5 w-5 mr-2" />
                    {participants.length}
                  </Button>
                  
                  <DropdownMenu>
                    <DropdownMenuTrigger asChild>
                      <Button variant="ghost" size="sm" className="text-white">
                        <MoreVertical className="h-5 w-5" />
                      </Button>
                    </DropdownMenuTrigger>
                    <DropdownMenuContent>
                      <DropdownMenuItem>
                        <Settings className="h-4 w-4 mr-2" />
                        Settings
                      </DropdownMenuItem>
                      <DropdownMenuItem>
                        <Share className="h-4 w-4 mr-2" />
                        Invite Others
                      </DropdownMenuItem>
                      <DropdownMenuItem>
                        <Camera className="h-4 w-4 mr-2" />
                        Record Session
                      </DropdownMenuItem>
                    </DropdownMenuContent>
                  </DropdownMenu>
                </div>
              </div>
            </div>
          </div>

          {/* Chat Sidebar */}
          {isChatOpen && (
            <div className="w-80 bg-white border-l border-gray-200 flex flex-col">
              <div className="p-4 border-b border-gray-200">
                <div className="flex items-center justify-between">
                  <h3 className="font-semibold text-gray-900">Live Chat</h3>
                  <Button
                    variant="ghost"
                    size="sm"
                    onClick={() => setIsChatOpen(false)}
                  >
                    ×
                  </Button>
                </div>
                {raisedHands > 0 && (
                  <div className="mt-2 p-2 bg-yellow-50 border border-yellow-200 rounded-lg">
                    <div className="flex items-center space-x-2">
                      <Hand className="h-4 w-4 text-yellow-600" />
                      <span className="text-sm text-yellow-800">
                        {raisedHands} student{raisedHands > 1 ? 's' : ''} raised hand{raisedHands > 1 ? 's' : ''}
                      </span>
                    </div>
                  </div>
                )}
              </div>
              
              <ScrollArea className="flex-1 p-4">
                <div className="space-y-4">
                  {chatMessages.map((message) => (
                    <div key={message.id} className={cn(
                      "flex space-x-2",
                      message.type === 'system' && "justify-center"
                    )}>
                      {message.type === 'message' && (
                        <Avatar className="h-6 w-6 mt-1">
                          <AvatarImage src={participants.find(p => p.id === message.participantId)?.avatar} />
                          <AvatarFallback className="text-xs">
                            {message.participantName.split(' ').map(n => n[0]).join('')}
                          </AvatarFallback>
                        </Avatar>
                      )}
                      
                      <div className={cn(
                        "flex-1",
                        message.type === 'system' && "text-center"
                      )}>
                        {message.type === 'system' ? (
                          <p className="text-xs text-gray-500 bg-gray-100 rounded-full px-3 py-1 inline-block">
                            {message.message}
                          </p>
                        ) : (
                          <>
                            <div className="flex items-center space-x-2 mb-1">
                              <span className="text-sm font-medium text-gray-900">
                                {message.participantName}
                              </span>
                              <span className="text-xs text-gray-500">{message.timestamp}</span>
                            </div>
                            <p className="text-sm text-gray-700">{message.message}</p>
                          </>
                        )}
                      </div>
                    </div>
                  ))}
                  <div ref={chatEndRef} />
                </div>
              </ScrollArea>
              
              <div className="p-4 border-t border-gray-200">
                <div className="flex space-x-2">
                  <Input
                    value={chatMessage}
                    onChange={(e) => setChatMessage(e.target.value)}
                    placeholder="Type a message..."
                    onKeyPress={(e) => e.key === 'Enter' && handleSendMessage()}
                    className="flex-1"
                  />
                  <Button onClick={handleSendMessage} size="sm">
                    <Send className="h-4 w-4" />
                  </Button>
                </div>
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  )
}