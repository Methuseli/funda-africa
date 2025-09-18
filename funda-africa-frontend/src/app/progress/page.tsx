"use client"

import React from 'react'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { Progress } from '@/components/ui/progress'
import { Calendar, Clock, Trophy, TrendingUp, Target, BookOpen } from 'lucide-react'
import Sidebar from '@/components/layout/sidebar'
import ProgressChart from '@/components/progress/progress-chart'
import { subjects } from '@/lib/data'

export default function ProgressPage() {
  const overallStats = {
    totalHours: 47,
    streakDays: 12,
    completedSubtopics: 23,
    totalSubtopics: 45,
    averageScore: 87,
    rank: 'Advanced Learner'
  }

  const weeklyGoals = [
    { goal: 'Complete 5 subtopics', progress: 80, completed: 4, target: 5 },
    { goal: 'Study 10 hours', progress: 65, completed: 6.5, target: 10 },
    { goal: 'Score 85%+ on quizzes', progress: 100, completed: 3, target: 3 }
  ]

  return (
    <div className="flex w-full">
      <Sidebar />
      
      <div className="flex-1 p-8">
        <div className="max-w-7xl mx-auto space-y-8">
          {/* Header */}
          <div>
            <h1 className="text-3xl font-bold text-gray-900">Your Learning Progress</h1>
            <p className="text-gray-600 mt-2">Track your achievements and continue your learning journey</p>
          </div>

          {/* Quick Stats */}
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
            <Card>
              <CardContent className="p-6">
                <div className="flex items-center justify-between">
                  <div>
                    <p className="text-sm font-medium text-gray-600">Study Hours</p>
                    <p className="text-2xl font-bold text-gray-900">{overallStats.totalHours}</p>
                  </div>
                  <Clock className="h-8 w-8 text-blue-600" />
                </div>
              </CardContent>
            </Card>
            
            <Card>
              <CardContent className="p-6">
                <div className="flex items-center justify-between">
                  <div>
                    <p className="text-sm font-medium text-gray-600">Study Streak</p>
                    <p className="text-2xl font-bold text-gray-900">{overallStats.streakDays} days</p>
                  </div>
                  <Calendar className="h-8 w-8 text-green-600" />
                </div>
              </CardContent>
            </Card>
            
            <Card>
              <CardContent className="p-6">
                <div className="flex items-center justify-between">
                  <div>
                    <p className="text-sm font-medium text-gray-600">Completion Rate</p>
                    <p className="text-2xl font-bold text-gray-900">
                      {Math.round((overallStats.completedSubtopics / overallStats.totalSubtopics) * 100)}%
                    </p>
                  </div>
                  <Trophy className="h-8 w-8 text-yellow-600" />
                </div>
              </CardContent>
            </Card>
            
            <Card>
              <CardContent className="p-6">
                <div className="flex items-center justify-between">
                  <div>
                    <p className="text-sm font-medium text-gray-600">Average Score</p>
                    <p className="text-2xl font-bold text-gray-900">{overallStats.averageScore}%</p>
                  </div>
                  <TrendingUp className="h-8 w-8 text-purple-600" />
                </div>
              </CardContent>
            </Card>
          </div>

          {/* Main Content Grid */}
          <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
            {/* Progress Charts */}
            <div className="lg:col-span-2">
              <ProgressChart />
            </div>

            {/* Weekly Goals */}
            <div className="space-y-6">
              <Card>
                <CardHeader>
                  <CardTitle className="flex items-center space-x-2">
                    <Target className="h-5 w-5" />
                    <span>Weekly Goals</span>
                  </CardTitle>
                  <CardDescription>Your progress towards this week's learning goals</CardDescription>
                </CardHeader>
                <CardContent className="space-y-4">
                  {weeklyGoals.map((goal, index) => (
                    <div key={index} className="space-y-2">
                      <div className="flex items-center justify-between text-sm">
                        <span className="font-medium">{goal.goal}</span>
                        <Badge variant={goal.progress === 100 ? "default" : "secondary"}>
                          {goal.completed}/{goal.target}
                        </Badge>
                      </div>
                      <Progress value={goal.progress} className="h-2" />
                      <p className="text-xs text-gray-600">
                        {goal.progress}% complete
                      </p>
                    </div>
                  ))}
                </CardContent>
              </Card>

              {/* Learner Rank */}
              <Card>
                <CardHeader>
                  <CardTitle>Current Rank</CardTitle>
                </CardHeader>
                <CardContent>
                  <div className="text-center space-y-4">
                    <div className="w-16 h-16 mx-auto bg-gradient-to-r from-purple-600 to-blue-600 rounded-full flex items-center justify-center">
                      <Trophy className="h-8 w-8 text-white" />
                    </div>
                    <div>
                      <h3 className="text-xl font-bold text-gray-900">{overallStats.rank}</h3>
                      <p className="text-sm text-gray-600">Keep learning to advance!</p>
                    </div>
                    <div className="space-y-2">
                      <div className="flex items-center justify-between text-xs">
                        <span>Progress to Expert</span>
                        <span>75%</span>
                      </div>
                      <Progress value={75} className="h-2" />
                    </div>
                  </div>
                </CardContent>
              </Card>
            </div>
          </div>

          {/* Subject Progress Detail */}
          <Card>
            <CardHeader>
              <CardTitle className="flex items-center space-x-2">
                <BookOpen className="h-5 w-5" />
                <span>Subject Progress Details</span>
              </CardTitle>
              <CardDescription>Detailed breakdown of your progress across all subjects</CardDescription>
            </CardHeader>
            <CardContent>
              <div className="space-y-6">
                {subjects.map((subject) => {
                  const subjectProgress = Math.floor(Math.random() * 60) + 20
                  const totalSubtopics = subject.topics.reduce((acc, topic) => acc + topic.subtopics.length, 0)
                  const completedSubtopics = Math.floor(totalSubtopics * (subjectProgress / 100))
                  const totalTime = subject.topics.reduce((acc, topic) => 
                    acc + topic.subtopics.reduce((topicAcc, subtopic) => topicAcc + subtopic.estimatedTime, 0), 0
                  )
                  
                  return (
                    <div key={subject.id} className="border rounded-lg p-6 space-y-4">
                      <div className="flex items-center justify-between">
                        <div className="flex items-center space-x-3">
                          <div className={`w-10 h-10 rounded-lg ${subject.color} flex items-center justify-center text-white`}>
                            <BookOpen className="h-5 w-5" />
                          </div>
                          <div>
                            <h3 className="text-lg font-semibold text-gray-900">{subject.title}</h3>
                            <p className="text-sm text-gray-600">{subject.description}</p>
                          </div>
                        </div>
                        <div className="text-right">
                          <p className="text-2xl font-bold text-gray-900">{subjectProgress}%</p>
                          <p className="text-sm text-gray-600">{completedSubtopics}/{totalSubtopics} completed</p>
                        </div>
                      </div>
                      
                      <Progress value={subjectProgress} className="h-3" />
                      
                      <div className="grid grid-cols-3 gap-4 text-sm">
                        <div className="text-center">
                          <p className="font-medium text-gray-900">{subject.topics.length}</p>
                          <p className="text-gray-600">Topics</p>
                        </div>
                        <div className="text-center">
                          <p className="font-medium text-gray-900">{Math.round(totalTime / 60)}h</p>
                          <p className="text-gray-600">Est. Time</p>
                        </div>
                        <div className="text-center">
                          <p className="font-medium text-gray-900">{Math.floor(Math.random() * 20) + 80}%</p>
                          <p className="text-gray-600">Avg. Score</p>
                        </div>
                      </div>
                    </div>
                  )
                })}
              </div>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  )
}