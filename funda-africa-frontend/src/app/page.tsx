"use client"

import React from 'react'
import Link from 'next/link'
import { ArrowRight, BookOpen, TrendingUp, Trophy, Clock } from 'lucide-react'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { subjects } from '@/lib/data'
import { useAppStore } from '@/lib/store'
import Sidebar from '@/components/layout/sidebar'
import ProgressChart from '@/components/progress/progress-chart'

const iconMap = {
  Calculator: '📐',
  Zap: '⚡',
  Code: '💻',
  BookOpen: '📚',
}

export default function Dashboard() {
  const { user } = useAppStore()

  const recentActivity = [
    {
      id: '1',
      title: 'Completed Vector Basics',
      subject: 'Mathematics',
      time: '2 hours ago',
      type: 'completion'
    },
    {
      id: '2',
      title: 'Started Newton\'s Laws',
      subject: 'Physics',
      time: '1 day ago',
      type: 'start'
    },
    {
      id: '3',
      title: 'Quiz: Sorting Algorithms - 95%',
      subject: 'Computer Science',
      time: '2 days ago',
      type: 'quiz'
    }
  ]

  const achievements = [
    { id: '1', title: 'First Steps', description: 'Complete your first subtopic', earned: true },
    { id: '2', title: 'Speed Learner', description: 'Complete 5 subtopics in one day', earned: true },
    { id: '3', title: 'Quiz Master', description: 'Score 90%+ on 3 quizzes', earned: false },
    { id: '4', title: 'Consistent Learner', description: 'Study for 7 consecutive days', earned: false },
  ]

  return (
    <div className="flex w-full">
      <Sidebar />
      
      <div className="flex-1 p-8">
        <div className="max-w-7xl mx-auto space-y-8">
          {/* Welcome Section */}
          <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4">
            <div>
              <h1 className="text-3xl font-bold text-gray-900">
                Welcome back, {user?.name || 'Learner'}! 👋
              </h1>
              <p className="text-gray-600 mt-1">
                Continue your learning journey and unlock your potential
              </p>
            </div>
            <Button size="lg" className="bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700">
              Continue Learning
              <ArrowRight className="ml-2 h-4 w-4" />
            </Button>
          </div>

          {/* Quick Stats */}
          <div className="grid grid-cols-1 md:grid-cols-4 gap-6">
            <Card>
              <CardContent className="p-6">
                <div className="flex items-center space-x-2">
                  <TrendingUp className="h-8 w-8 text-blue-600" />
                  <div>
                    <p className="text-sm font-medium text-gray-600">Overall Progress</p>
                    <p className="text-2xl font-bold text-gray-900">73%</p>
                  </div>
                </div>
              </CardContent>
            </Card>
            
            <Card>
              <CardContent className="p-6">
                <div className="flex items-center space-x-2">
                  <BookOpen className="h-8 w-8 text-green-600" />
                  <div>
                    <p className="text-sm font-medium text-gray-600">Subjects</p>
                    <p className="text-2xl font-bold text-gray-900">{subjects.length}</p>
                  </div>
                </div>
              </CardContent>
            </Card>
            
            <Card>
              <CardContent className="p-6">
                <div className="flex items-center space-x-2">
                  <Trophy className="h-8 w-8 text-yellow-600" />
                  <div>
                    <p className="text-sm font-medium text-gray-600">Achievements</p>
                    <p className="text-2xl font-bold text-gray-900">
                      {achievements.filter(a => a.earned).length}
                    </p>
                  </div>
                </div>
              </CardContent>
            </Card>
            
            <Card>
              <CardContent className="p-6">
                <div className="flex items-center space-x-2">
                  <Clock className="h-8 w-8 text-purple-600" />
                  <div>
                    <p className="text-sm font-medium text-gray-600">Study Streak</p>
                    <p className="text-2xl font-bold text-gray-900">5 days</p>
                  </div>
                </div>
              </CardContent>
            </Card>
          </div>

          {/* Main Content Grid */}
          <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
            {/* Subjects Overview */}
            <div className="lg:col-span-2 space-y-6">
              <h2 className="text-2xl font-semibold text-gray-900">Your Subjects</h2>
              <div className="grid gap-6">
                {subjects.map((subject) => {
                  const progress = Math.floor(Math.random() * 40) + 30 // Mock progress
                  return (
                    <Card key={subject.id} className="hover:shadow-lg transition-shadow duration-200">
                      <CardContent className="p-6">
                        <div className="flex items-center justify-between">
                          <div className="flex items-center space-x-4">
                            <div className={`w-12 h-12 rounded-lg ${subject.color} flex items-center justify-center text-2xl`}>
                              {iconMap[subject.icon as keyof typeof iconMap]}
                            </div>
                            <div>
                              <h3 className="text-lg font-semibold text-gray-900">{subject.title}</h3>
                              <p className="text-gray-600">{subject.description}</p>
                              <p className="text-sm text-gray-500 mt-1">
                                {subject.topics.length} topics • 
                                {subject.topics.reduce((acc, topic) => acc + topic.subtopics.length, 0)} subtopics
                              </p>
                            </div>
                          </div>
                          <div className="text-right">
                            <p className="text-2xl font-bold text-gray-900">{progress}%</p>
                            <Link href={`/subjects/${subject.id}`}>
                              <Button variant="outline" size="sm" className="mt-2">
                                Continue
                                <ArrowRight className="ml-1 h-3 w-3" />
                              </Button>
                            </Link>
                          </div>
                        </div>
                      </CardContent>
                    </Card>
                  )
                })}
              </div>
            </div>

            {/* Activity & Achievements Sidebar */}
            <div className="space-y-6">
              {/* Recent Activity */}
              <Card>
                <CardHeader>
                  <CardTitle>Recent Activity</CardTitle>
                </CardHeader>
                <CardContent className="space-y-4">
                  {recentActivity.map((activity) => (
                    <div key={activity.id} className="flex items-start space-x-3">
                      <div className="w-2 h-2 bg-blue-600 rounded-full mt-2"></div>
                      <div className="flex-1">
                        <p className="text-sm font-medium text-gray-900">{activity.title}</p>
                        <p className="text-xs text-gray-600">{activity.subject}</p>
                        <p className="text-xs text-gray-400">{activity.time}</p>
                      </div>
                    </div>
                  ))}
                </CardContent>
              </Card>

              {/* Recent Achievements */}
              <Card>
                <CardHeader>
                  <CardTitle>Achievements</CardTitle>
                </CardHeader>
                <CardContent className="space-y-3">
                  {achievements.slice(0, 3).map((achievement) => (
                    <div key={achievement.id} className="flex items-center space-x-3">
                      <Trophy className={`h-6 w-6 ${achievement.earned ? 'text-yellow-500' : 'text-gray-300'}`} />
                      <div className="flex-1">
                        <p className={`text-sm font-medium ${achievement.earned ? 'text-gray-900' : 'text-gray-500'}`}>
                          {achievement.title}
                        </p>
                        <p className="text-xs text-gray-600">{achievement.description}</p>
                      </div>
                      {achievement.earned && <Badge variant="secondary" className="text-xs">Earned</Badge>}
                    </div>
                  ))}
                  <Link href="/achievements">
                    <Button variant="ghost" size="sm" className="w-full mt-2">
                      View All Achievements
                    </Button>
                  </Link>
                </CardContent>
              </Card>
            </div>
          </div>

          {/* Progress Charts */}
          <div>
            <h2 className="text-2xl font-semibold text-gray-900 mb-6">Progress Analytics</h2>
            <ProgressChart />
          </div>
        </div>
      </div>
    </div>
  )
}