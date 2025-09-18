import { create } from 'zustand'
import { persist } from 'zustand/middleware'

export interface User {
  id: string
  name: string
  email: string
  avatar?: string
  joinDate: string
}

export interface SubtopicProgress {
  subtopicId: string
  completed: boolean
  progress: number
  timeSpent: number
  lastAccessed: string
  quizScores?: number[]
}

export interface TopicProgress {
  topicId: string
  progress: number
  subtopics: SubtopicProgress[]
}

export interface SubjectProgress {
  subjectId: string
  progress: number
  topics: TopicProgress[]
}

export interface ChatMessage {
  id: string
  content: string
  sender: 'user' | 'ai'
  timestamp: string
  subtopicId?: string
}

interface AppState {
  user: User | null
  progress: SubjectProgress[]
  chatHistory: ChatMessage[]
  notifications: any[]
  isChatOpen: boolean
  currentSubtopic: string | null
  
  // Actions
  setUser: (user: User) => void
  updateSubtopicProgress: (subtopicId: string, progress: Partial<SubtopicProgress>) => void
  addChatMessage: (message: ChatMessage) => void
  toggleChat: () => void
  setCurrentSubtopic: (subtopicId: string | null) => void
  addNotification: (notification: any) => void
  markNotificationAsRead: (id: string) => void
}

export const useAppStore = create<AppState>()(
  persist(
    (set, get) => ({
      user: {
        id: '1',
        name: 'Alex Johnson',
        email: 'alex@example.com',
        avatar: 'https://images.pexels.com/photos/614810/pexels-photo-614810.jpeg?auto=compress&cs=tinysrgb&w=400',
        joinDate: '2024-01-15'
      },
      progress: [],
      chatHistory: [],
      notifications: [
        {
          id: '1',
          title: 'New lesson available',
          message: 'Advanced Calculus: Integration by Parts is now available',
          timestamp: new Date().toISOString(),
          read: false,
          type: 'lesson'
        },
        {
          id: '2',
          title: 'Achievement unlocked',
          message: 'You\'ve completed 5 subtopics in Linear Algebra!',
          timestamp: new Date(Date.now() - 3600000).toISOString(),
          read: false,
          type: 'achievement'
        }
      ],
      isChatOpen: false,
      currentSubtopic: null,

      setUser: (user) => set({ user }),
      
      updateSubtopicProgress: (subtopicId, progressUpdate) => set((state) => {
        const newProgress = [...state.progress]
        // Implementation for updating nested progress structure
        return { progress: newProgress }
      }),

      addChatMessage: (message) => set((state) => ({
        chatHistory: [...state.chatHistory, message]
      })),

      toggleChat: () => set((state) => ({
        isChatOpen: !state.isChatOpen
      })),

      setCurrentSubtopic: (subtopicId) => set({ currentSubtopic: subtopicId }),

      addNotification: (notification) => set((state) => ({
        notifications: [notification, ...state.notifications]
      })),

      markNotificationAsRead: (id) => set((state) => ({
        notifications: state.notifications.map(n => 
          n.id === id ? { ...n, read: true } : n
        )
      }))
    }),
    {
      name: 'funda-africa-store',
    }
  )
)