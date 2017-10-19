## CircleCI configurations.

A workflow with two jobs ( ``Build -> Deploy``) was configured. ``Deploy`` is only triggered if ``Build`` finishes successfully.

Later, it is possible to configure ``Deploy`` to only be triggered if a specific condition is met. At the moment, it will trigger if something is pushed to the ``develop`` branch.

