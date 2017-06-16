/**
 * Copyright 2015 The Go Authors. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package org.golang.mobile

import org.gradle.api.Task
import org.gradle.api.artifacts.PublishArtifact
import org.gradle.api.tasks.TaskDependency

import java.io.File
import java.util.Date

/**
 * custom implementation of PublishArtifact for published AAR
 */
class AARPublishArtifact(
        private val name: String,
        private val classifier: String,
        private val task: OutputFileTask) : PublishArtifact {
    private val taskDependency: TaskDependency

    private class DefaultTaskDependency internal constructor(task: Task) : TaskDependency {

        private val tasks: Set<Task>

        init {
            this.tasks = setOf(task)
        }

        override fun getDependencies(task: Task): Set<Task> {
            return tasks
        }
    }

    init {
        this.taskDependency = DefaultTaskDependency(task as Task)
    }


    override fun getName(): String {
        return name
    }

    override fun getExtension(): String {
        return "aar"
    }

    override fun getType(): String {
        return "aar"
    }

    override fun getClassifier(): String {
        return classifier
    }

    override fun getFile(): File {
        return task.outputFile
    }

    override fun getDate(): Date? {
        return null
    }

    override fun getBuildDependencies(): TaskDependency {
        return taskDependency
    }
}
