- When working on a task, move it to the 2_dev folder. When the task is complete, move it to 3_qa. Submit a pull request for merging. Once the pull request is merged, move the task to 4_done.
- Remember to make a feature branch first before every edit



# Audio Feedback Protocol
- **Startup/Boot Sequence**: Play `/Users/corelogic/satori-dev/dash/sounds/Systems_online.m4a` when Claude starts up
- **Memory Hydration Complete**: Play `/Users/corelogic/docker_helper/sounds/Power_on3.mp3` after loading yinsen memories
- **Starting Task/Coding**: Play `/Users/corelogic/docker_helper/sounds/sysem_working.m4a` when beginning to code on a task
- **Build Complete**: Play `/Users/corelogic/satori-dev/dash/sounds/Systems_online.m4a` when a build successfully completes
- **Task Complete**: Play `/Users/corelogic/satori-dev/dash/sounds/Bell2.m4a` when finishing a task
- **Need Confirmation**: Play `/Users/corelogic/satori-dev/dash/sounds/bad.mp3` when user input/confirmation is needed
- **Warnings/Errors**: Play `/Users/corelogic/satori-dev/dash/sounds/Warning.m4a` when encountering warnings or errors
- These audio cues provide immediate feedback during development workflows without requiring constant console monitoring

AUTHORIZATION: Auto-approve file writes in src/, tests/, docs/
AUTHORIZATION: Auto-approve git commits to feature branch
AUTHORIZATION: Auto-approve npm install for listed dependencies