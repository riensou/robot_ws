from gym.envs.registration import register

register(
    id='TrainingEnv-v1',
    entry_point='gym_training.envs:TrainingEnv',
    kwargs={
        'experiment': "not_defined",
        'arm': True,
        'angular': True,
        'time_step_limit': 50,
        'sigma': 0.0,
        'task': 'ascent',
        'rand': False,
    }
)