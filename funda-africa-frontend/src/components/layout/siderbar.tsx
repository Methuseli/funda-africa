"use client"

import React from 'react'
import Link from 'next/link'
import { usePathname } from 'next/navigation'
import { 
  Home, 
  BookOpen, 
  TrendingUp, 
  Trophy, 
  Settings,
  ChevronRight,
  Calculator,
  Zap,
  Code
} from 'lucide-react'
import { cn } from '@/lib/utils'
import { subjects } from '@/lib/data'
import { Button } from '@/components/ui/button'
import {
  Collapsible,
  CollapsibleContent,
  CollapsibleTrigger,
} from '@/components/ui/collapsible'

const navigation = [
  { name: 'Dashboard', href: '/', icon: Home },
  { name: 'Progress', href: '/progress', icon: TrendingUp },
  { name: 'Achievements', href: '/achievements', icon: Trophy },
  { name: 'Settings', href: '/settings', icon: Settings },
]

const iconMap = {
  Calculator: Calculator,
  Zap: Zap,
  Code: Code,
  BookOpen: BookOpen,
}

export default function Sidebar() {
  const pathname = usePathname()

  return (
    <div className="w-64 h-full bg-white border-r border-gray-200 overflow-y-auto">
      <div className="p-6">
        {/* Main Navigation */}
        <nav className="space-y-2 mb-8">
          {navigation.map((item) => {
            const Icon = item.icon
            return (
              <Link
                key={item.name}
                href={item.href}
                className={cn(
                  'flex items-center px-3 py-2 text-sm font-medium rounded-lg transition-colors',
                  pathname === item.href
                    ? 'bg-blue-50 text-blue-700 border border-blue-200'
                    : 'text-gray-600 hover:text-gray-900 hover:bg-gray-50'
                )}
              >
                <Icon className="mr-3 h-5 w-5" />
                {item.name}
              </Link>
            )
          })}
        </nav>

        {/* Subjects */}
        <div>
          <h3 className="px-3 mb-4 text-xs font-semibold text-gray-500 uppercase tracking-wider">
            Subjects
          </h3>
          <div className="space-y-2">
            {subjects.map((subject) => {
              const Icon = iconMap[subject.icon as keyof typeof iconMap] || BookOpen
              const isActive = pathname.includes(`/subjects/${subject.id}`)
              
              return (
                <Collapsible key={subject.id} defaultOpen={isActive}>
                  <CollapsibleTrigger asChild>
                    <Button
                      variant="ghost"
                      className={cn(
                        'w-full justify-start px-3 py-2 text-sm font-medium',
                        isActive ? 'bg-gray-50 text-gray-900' : 'text-gray-600 hover:text-gray-900'
                      )}
                    >
                      <Icon className="mr-3 h-4 w-4" />
                      <span className="truncate">{subject.title}</span>
                      <ChevronRight className="ml-auto h-4 w-4 transition-transform duration-200" />
                    </Button>
                  </CollapsibleTrigger>
                  <CollapsibleContent className="space-y-1">
                    {subject.topics.map((topic) => (
                      <Link
                        key={topic.id}
                        href={`/subjects/${subject.id}/${topic.id}`}
                        className={cn(
                          'block px-9 py-1.5 text-sm text-gray-600 hover:text-gray-900 hover:bg-gray-50 rounded-md transition-colors',
                          pathname.includes(`/${topic.id}`) && 'text-blue-600 bg-blue-50'
                        )}
                      >
                        {topic.title}
                      </Link>
                    ))}
                  </CollapsibleContent>
                </Collapsible>
              )
            })}
          </div>
        </div>
      </div>
    </div>
  )
}